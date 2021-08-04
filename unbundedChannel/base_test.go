package unbundedChannel

import (
	"sync"
	"testing"
)

func TestMakeInfinite(t *testing.T) {
	in, out := MakeInfinite()
	lastVal := -1
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for item := range out {
			vi := item.(int)
			if lastVal+1 != vi {
				t.Errorf("unexpected value")
			}
			lastVal = vi
		}
		wg.Done()
	}()

	for i := 0; i < 100; i++ {
		in <- i
	}
	close(in)
	wg.Wait()
	if lastVal != 99 {
		t.Errorf("last one received was :%v\n", lastVal)
	}
}
