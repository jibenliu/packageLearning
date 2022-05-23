package concurrency

import (
	"sync"
	"testing"
)

func BenchmarkBaseCounter(b *testing.B) {
	counter := NewCounter()
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		if i%3 == 0 {
			go func(counter *Counter) {
				counter.Read()
				wg.Done()
			}(counter)
		} else if i%3 == 1 {
			go func(counter *Counter) {
				counter.Add(1)
				counter.Read()
				wg.Done()
			}(counter)
		} else {
			go func(counter *Counter) {
				counter.Add(1)
				wg.Done()
			}(counter)
		}
	}

	wg.Wait()

	if counter.Read() != 66 {
		b.Errorf("counter should be %d and was %d", 66, counter.Read())
	}
}

func BenchmarkMutexCounter(b *testing.B) {
	counter := NewMutexCounter()
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		if i%3 == 0 {
			go func(counter *MutexCounter) {
				counter.Read()
				wg.Done()
			}(counter)
		} else if i%3 == 1 {
			go func(counter *MutexCounter) {
				counter.Add(1)
				counter.Read()
				wg.Done()
			}(counter)
		} else {
			go func(counter *MutexCounter) {
				counter.Add(1)
				wg.Done()
			}(counter)
		}
	}

	wg.Wait()

	if counter.Read() != 66 {
		b.Errorf("mutex counter should be %d and was %d", 66, counter.Read())
	}
}

func BenchmarkChanCounter(b *testing.B) {
	counter := NewChannelCounter()
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		if i%3 == 0 {
			go func(counter *ChannelCounter) {
				counter.Read()
				wg.Done()
			}(counter)
		} else if i%3 == 1 {
			go func(counter *ChannelCounter) {
				counter.Add(1)
				counter.Read()
				wg.Done()
			}(counter)
		} else {
			go func(counter *ChannelCounter) {
				counter.Add(1)
				wg.Done()
			}(counter)
		}
	}

	wg.Wait()

	if counter.Read() != 66 {
		b.Errorf("channel counter should be %d and was %d", 66, counter.Read())
	}
}

func BenchmarkAtomicCounter(b *testing.B) {
	counter := NewAtomicCounter()
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		if i%3 == 0 {
			go func(counter *AtomicCounter) {
				counter.Read()
				wg.Done()
			}(counter)
		} else if i%3 == 1 {
			go func(counter *AtomicCounter) {
				counter.Add(1)
				counter.Read()
				wg.Done()
			}(counter)
		} else {
			go func(counter *AtomicCounter) {
				counter.Add(1)
				wg.Done()
			}(counter)
		}
	}

	wg.Wait()

	if counter.Read() != 66 {
		b.Errorf("atomic counter should be %d and was %d", 66, counter.Read())
	}
}

func BenchmarkAtomic2Counter(b *testing.B) {
	counter := NewAtomic2Counter()
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		if i%3 == 0 {
			go func(counter *Atomic2Counter) {
				counter.Read()
				wg.Done()
			}(counter)
		} else if i%3 == 1 {
			go func(counter *Atomic2Counter) {
				counter.Add(1)
				counter.Read()
				wg.Done()
			}(counter)
		} else {
			go func(counter *Atomic2Counter) {
				counter.Add(1)
				wg.Done()
			}(counter)
		}
	}

	wg.Wait()

	if counter.Read() != 66 {
		b.Errorf("atomic2 counter should be %d and was %d", 66, counter.Read())
	}
}
