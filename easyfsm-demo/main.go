package main

import (
	"fmt"
	"github.com/wuqinqiang/easyfsm"
	"time"
)

var (
	// 业务
	businessName easyfsm.BusinessName = "order"

	// 对应状态
	initState easyfsm.State = 1 // 初始化
	paidState easyfsm.State = 2 // 已付款
	canceled  easyfsm.State = 3 // 已取消

	//对应事件
	paymentOrderEventName easyfsm.EventName = "paymentOrderEventName"
	cancelOrderEventName  easyfsm.EventName = "cancelOrderEventName"
)

type (
	orderParam struct {
		OrderNo string
	}
)

var (
	_ easyfsm.EventObserver = (*NotifyExample)(nil)
	_ easyfsm.EventHook     = (*HookExample)(nil)
)

type (
	NotifyExample struct {
	}
	HookExample struct {
	}
)

func (h HookExample) Before(opt *easyfsm.Param) {
	fmt.Println("事件执行前")
}

func (h HookExample) After(opt easyfsm.Param, state easyfsm.State, err error) {
	fmt.Println("事件执行后")
}

func (o NotifyExample) Receive(opt *easyfsm.Param) {
	fmt.Println("接收到事件变动,发送消息")
}

func init() {
	// 支付订单事件
	entity := easyfsm.NewEventEntity(paymentOrderEventName,
		func(opt *easyfsm.Param) (easyfsm.State, error) {
			param, ok := opt.Data.(orderParam)
			if !ok {
				panic("param err")
			}
			fmt.Printf("param:%+v\n", param)
			// 处理核心业务
			return paidState, nil
		}, easyfsm.WithHook(HookExample{}), easyfsm.WithObservers(NotifyExample{}))

	// 取消订单事件
	cancelEntity := easyfsm.NewEventEntity(cancelOrderEventName,
		func(opt *easyfsm.Param) (easyfsm.State, error) {
			// 处理核心业务
			param, ok := opt.Data.(orderParam)
			if !ok {
				panic("param err")
			}
			fmt.Printf("param:%+v\n", param)
			return canceled, nil
		}, easyfsm.WithHook(HookExample{}))

	// 注册订单状态机，初始状态为未支付，后续状态可以为支付和取消支付
	easyfsm.RegisterStateMachine(businessName,
		initState,
		entity, cancelEntity)
	// 注册订单状态机，初始状态为已支付，后续状态可以为取消支付
	easyfsm.RegisterStateMachine(businessName,
		paidState,
		cancelEntity)
}

func main() {

	// 第一步根据业务，以及当前状态生成fsm
	fsm := easyfsm.NewFSM(businessName, initState)

	// 第二步 调用支付完成具体
	currentState, err := fsm.Call(paymentOrderEventName,
		easyfsm.WithData(orderParam{OrderNo: "wuqinqiang050@gmail.com"}))

	fmt.Printf("[Success]call paymentOrderEventName err:%v\n", err)
	fmt.Printf("[Success]call paymentOrderEventName state:%v\n", currentState)

	// 第三步 调用取消支付具体
	currentState, err = fsm.Call(cancelOrderEventName,
		easyfsm.WithData(orderParam{OrderNo: "wuqinqiang050@gmail.com"}))

	fmt.Printf("[Success]call paymentOrderEventName err:%v\n", err)
	fmt.Printf("[Success]call paymentOrderEventName state:%v\n", currentState)

	time.Sleep(5 * time.Second)
}
