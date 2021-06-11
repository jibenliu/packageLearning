package serviceBreaker

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

var (
	ErrStateOpen    = errors.New("service breaker is open")
	ErrTooManyCalls = errors.New("service breaker is halfOpen, too many calls")
)

type ServiceBreaker struct {
	mu               sync.Mutex
	name             string
	state            State
	windowInterval   time.Duration
	metrics          Metrics
	tripStrategyFunc TripStrategyFunc
	halfMaxCalls     uint64
	stateOpenTime    time.Time
	sleepTimeout     time.Duration
	stateChangeHook  func(name string, fromState, toState State)
}

type State int

const (
	StateClosed State = iota
	StateOpen
	StateHalfOpen
)

const (
	ConsecutiveFailTrip = iota + 1
	FailTrip
	FailRateTrip
)

func (s State) String() string {
	switch s {
	case StateClosed:
		return "closed"
	case StateHalfOpen:
		return "half-open"
	case StateOpen:
		return "open"
	default:
		return fmt.Sprintf("unknown state:%d", s)
	}
}

type TripStrategyOption struct {
	Strategy                 uint
	ConsecutiveFailThreshold uint64
	FailThreshold            uint64
	FailRate                 float64
	MinCall                  uint64
}
type Option struct {
	Name            string
	WindowInterval  time.Duration
	HalfMaxCalls    uint64
	SleepTimeout    time.Duration
	StateChangeHook func(name string, fromState State, toState State)
	TripStrategy    TripStrategyOption
}
type Metrics struct {
	WindowBatch        uint64
	WindowTimeStart    time.Time
	CountAll           uint64
	CountSuccess       uint64
	CountFail          uint64
	ConsecutiveSuccess uint64
	ConsecutiveFail    uint64
}

func (m *Metrics) NewBatch() {
	m.WindowBatch++
}
func (m *Metrics) OnCall() {
	m.CountAll++
}
func (m *Metrics) OnSuccess() {
	m.CountSuccess++
	m.ConsecutiveSuccess++
	m.ConsecutiveFail = 0
}
func (m *Metrics) OnFail() {
	m.CountFail++
	m.ConsecutiveFail++
	m.ConsecutiveSuccess = 0
}
func (m *Metrics) OnReset() {
	m.CountAll = 0
	m.CountSuccess = 0
	m.CountFail = 0
	m.ConsecutiveSuccess = 0
	m.ConsecutiveFail = 0
}

// TripStrategyFunc when error occur, determine whether the breaker should be opened.
type TripStrategyFunc func(Metrics) bool

// ConsecutiveFailTripFunc according to consecutive fail
func ConsecutiveFailTripFunc(threshold uint64) TripStrategyFunc {
	return func(m Metrics) bool {
		return m.ConsecutiveFail >= threshold
	}
}

// FailTripFunc according to fail
func FailTripFunc(threshold uint64) TripStrategyFunc {
	return func(m Metrics) bool {
		return m.CountFail >= threshold
	}
}

// FailRateTripFunc according to fail rate
func FailRateTripFunc(rate float64, minCalls uint64) TripStrategyFunc {
	return func(m Metrics) bool {
		var currRate float64
		if m.CountAll != 0 {
			currRate = float64(m.CountFail) / float64(m.CountAll)
		}
		return m.CountAll >= minCalls && currRate >= rate
	}
}

// ChooseTrip choose trip
func ChooseTrip(op *TripStrategyOption) TripStrategyFunc {
	switch op.Strategy {
	case ConsecutiveFailTrip:
		return ConsecutiveFailTripFunc(op.ConsecutiveFailThreshold)
	case FailTrip:
		return FailTripFunc(op.FailThreshold)
	case FailRateTrip:
		fallthrough
	default:
		return FailRateTripFunc(op.FailRate, op.MinCall)
	}
}

func NewServiceBreaker(op Option) (*ServiceBreaker, error) {
	if op.WindowInterval <= 0 || op.HalfMaxCalls <= 0 || op.SleepTimeout <= 0 {
		return nil, errors.New("incomplete options")
	}
	breaker := new(ServiceBreaker)
	breaker.name = op.Name
	breaker.windowInterval = op.WindowInterval
	breaker.halfMaxCalls = op.HalfMaxCalls
	breaker.sleepTimeout = op.SleepTimeout
	breaker.stateChangeHook = op.StateChangeHook
	breaker.tripStrategyFunc = ChooseTrip(&op.TripStrategy)
	breaker.nextWindow(time.Now())
	return breaker, nil
}

func (breaker *ServiceBreaker) onSuccess(now time.Time) {
	breaker.metrics.OnSuccess()
	if breaker.state == StateHalfOpen && breaker.metrics.ConsecutiveSuccess >= breaker.halfMaxCalls {
		breaker.changeState(StateClosed, now)
	}
}

func (breaker *ServiceBreaker) changeState(state State, now time.Time) {
	if breaker.state == state {
		return
	}
	prevState := breaker.state
	breaker.state = state
	//goto next window,reset metrics
	breaker.nextWindow(time.Now())
	//record open time
	if state == StateOpen {
		breaker.stateOpenTime = now
	}
	//callback hook
	if breaker.stateChangeHook != nil {
		breaker.stateChangeHook(breaker.name, prevState, state)
	}
}

func (breaker *ServiceBreaker) nextWindow(now time.Time) {
	breaker.metrics.NewBatch()
	breaker.metrics.OnReset() //clear count num
	var zero time.Time
	switch breaker.state {
	case StateClosed:
		if breaker.windowInterval == 0 {
			breaker.metrics.WindowTimeStart = zero
		} else {
			breaker.metrics.WindowTimeStart = now.Add(breaker.windowInterval)
		}
	case StateOpen:
		breaker.metrics.WindowTimeStart = now.Add(breaker.sleepTimeout)
	default: //halfOpen
		breaker.metrics.WindowTimeStart = zero //halfOpen no window
	}
}

func (breaker *ServiceBreaker) Call(exec func() (interface{}, error)) (interface{}, error) {
	//before call
	err := breaker.beforeCall()
	if err != nil {
		return nil, err
	}
	//if panic occur
	defer func() {
		err := recover()
		if err != nil {
			breaker.afterCall(false)
			panic(err)
		}
	}()
	//call
	breaker.metrics.OnCall()
	result, err := exec()
	//after call
	breaker.afterCall(err == nil)
	return result, err
}

func (breaker *ServiceBreaker) beforeCall() error {
	breaker.mu.Lock()
	defer breaker.mu.Unlock()
	now := time.Now()
	switch breaker.state {
	case StateOpen:
		//after sleep timeout, can retry
		if breaker.stateOpenTime.Add(breaker.sleepTimeout).Before(now) {
			log.Printf("%s 熔断过冷却期，尝试半开\n", breaker.name)
			breaker.changeState(StateHalfOpen, now)
			return nil
		}
		log.Printf("%s 熔断打开，请求被阻止\n", breaker.name)
		return ErrStateOpen
	case StateHalfOpen:
		if breaker.metrics.CountAll >= breaker.halfMaxCalls {
			log.Printf("%s 熔断半开，请求过多被阻止\n", breaker.name)
			return ErrTooManyCalls
		}
	default: //Closed
		if !breaker.metrics.WindowTimeStart.IsZero() && breaker.metrics.WindowTimeStart.Before(now) {
			breaker.nextWindow(now)
			return nil
		}
	}
	return nil
}

func (breaker *ServiceBreaker) afterCall(success bool) {
	breaker.mu.Lock()
	defer breaker.mu.Unlock()
	if success {
		breaker.onSuccess(time.Now())
	} else {
		breaker.onFail(time.Now())
	}
}

func (breaker *ServiceBreaker) onFail(now time.Time) {
	breaker.metrics.OnFail()
	switch breaker.state {
	case StateClosed:
		if breaker.tripStrategyFunc(breaker.metrics) {
			breaker.changeState(StateOpen, now)
		}
	case StateHalfOpen:
		breaker.changeState(StateOpen, now)

	}
}
