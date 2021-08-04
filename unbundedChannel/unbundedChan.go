package unbundedChannel

type T interface{}

type UnboundedChan struct {
	In     chan<- T
	Out    <-chan T
	buffer []T
}

func (c UnboundedChan) Len() int {
	return len(c.buffer) + len(c.Out)
}

func (c UnboundedChan) BufLen() int {
	return len(c.buffer)
}

func NewUnboundedChan(initCapacity int) UnboundedChan {
	// 创建三个字段和无限缓存的chan类型
	in := make(chan T, initCapacity)
	out := make(chan T, initCapacity)
	ch := UnboundedChan{
		In:     in,
		Out:    out,
		buffer: make([]T, initCapacity),
	}

	// 通过一个goroutine,不断地从in中读取出来数据，放入到out或者buffer中
	go func() {
		defer close(out) // in关闭，数据读取完后也把out关闭
	loop:
		for {
			val, ok := <-in
			if !ok { // 如果in已经被closed, 退出loop
				break loop
			}
			// 否则尝试把从in中读取出来的数据放入到out中
			select {
			case out <- val: //放入成功，说明out刚才还没有满，buffer中也没有额外的数据待处理，所以回到loop开始
				continue
			default:
			}

			//如果out已经满了 需要把数据放入到缓存中
			ch.buffer = append(ch.buffer, val)
			// 处理缓存，一直尝试把缓存中的数据放入到out,直到缓存中没有数据了,
			// 为了避免阻塞住in channel,还要尝试从in中读取数据，因为这个时候out是满的，所以就直接把数据放入到缓存中
			for len(ch.buffer) > 0 {
				select {
				case val, ok := <-in: // 从in读取数据，放入到缓存中，如果in被closed, 退出loop
					if !ok {
						break loop
					}
					ch.buffer = append(ch.buffer, val)
				case out <- ch.buffer[0]: // 把缓存中最老的数据放入到out中，并移出第一个元素
					ch.buffer = ch.buffer[1:]
					if len(ch.buffer) == 0 { // 避免内存泄露. 如果缓存处理完了，恢复成原始的状态
						ch.buffer = make([]T, 0, initCapacity)
					}
				}
			}
		}
		// in被关闭，退出loop后，buffer中可能还有未处理的数据，需要把它们塞入到out中
		// 这个逻辑叫做"drain"。
		// 这一段逻辑处理完后，就可以把out关闭掉了
		for len(ch.buffer) > 0 {
			out <- ch.buffer[0]
			ch.buffer = ch.buffer[1:]
		}
	}()
	return ch
}
