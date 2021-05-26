package unbundedChannel

func MakeInfinite() (chan<- interface{}, <-chan interface{}) {
	in := make(chan interface{})  //写入数据
	out := make(chan interface{}) //输出数据

	go func() {
		var inQueue []interface{}
		//-------------方案2区块开始------
		outCh := func() chan interface{} {
			if len(inQueue) == 0 {
				return nil
			}
			return out
		}
		curlVal := func() interface{} {
			if len(inQueue) == 0 {
				return nil
			}
			return inQueue[0]
		}
		//-------------方案2区块结束------
		//loop:
		//for { //方案1
		for len(inQueue) > 0 || in != nil {
			select {
			case v, ok := <-in:
				if !ok { //写入通道关闭
					//break loop
					in = nil
				} else { //把从in拿出的数据放入到切片中
					inQueue = append(inQueue, v)
				}
			//case out <- inQueue[0]: //从切片中拿数据写入到out通道 //方案1
			case outCh() <- curlVal():
				inQueue = inQueue[1:]
			}
		}
		close(out)
	}()
	return in, out
}
