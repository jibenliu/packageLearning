package orDone

import (
	"reflect"
)

// refer : https://mp.weixin.qq.com/s/F5H0C--diuuPTcrZhSC0cA

// Or 传入多个并发chan，返回是否结束的 orDone chan
func Or(channels ...<-chan interface{}) <-chan interface{} {
	// 只有零个或者1个chan
	switch len(channels) {
	case 0:
		// 返回nil， 让读取阻塞等待
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		// 返回时利用close做结束信号的广播
		defer close(orDone)

		// 利用select监听第一个chan的返回
		switch len(channels) {
		case 2: // 直接select
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default: // 二分法递归处理
			m := len(channels) / 2
			select {
			case <-Or(channels[:m]...):
			case <-Or(channels[m:]...):
			}
		}
	}()

	return orDone
}

// SelectCase ----------------------------------------------方法二分割线------------------------

func OrInReflect(channels ...<-chan interface{}) <-chan interface{} {
	// 只有0个或者1个
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		// 利用反射构建SelectCase，这里是读取
		var cases []reflect.SelectCase
		for _, c := range channels {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}

		// 随机选择一个可用的case
		reflect.Select(cases)
	}()

	return orDone
}

// -------------------------------------------测试分割线-----------------------------------------------

func repeat(
	done <-chan interface{},
	// 外部传入done控制是否结束
	values ...interface{},
) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		// 返回时释放
		defer close(valueStream)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valueStream <- v:
				}
			}
		}
	}()
	return valueStream
}
