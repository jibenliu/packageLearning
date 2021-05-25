package main

// refer : https://mp.weixin.qq.com/s?__biz=MzAxMTA4Njc0OQ==&mid=2651439856&idx=3&sn=284dfa73ee1344ab4e9a8184450738c9&chksm=80bb1c02b7cc951445b9a46475a9ff8910c459e2118b811336143b9757b8f747096c86e6e619&xtrack=1&scene=0&subscene=91&sessionid=1591150816&clicktime=1591151545&enterid=1591151545&ascene=7&devicetype=android-26&version=27000f37&nettype=WIFI&abtest_cookie=AAACAA%3D%3D&lang=zh_CN&exportkey=ATvKuYRQc1DgJCMHO1BOKd8%3D&pass_ticket=2NmgE2PrcsJOWTEXgP7bBvwi31Vl0gTPnt9zX8FpcpKbB8Js1YfKvoUGXmKf%2BEzw&wx_header=1

// once是一个简单而强大的原语，可确保在并行程序中一个函数仅执行一次。
type Once chan struct{}

func NewOnce() Once {
	o := make(Once, 1)
	// 只允许一个goroutine接收，其他goroutine会被阻塞住
	o <- struct{}{}
	return o
}

func (o Once) Do(f func()) {
	_, ok := <-o
	if !ok {
		// Channel已经被关闭
		// 证明f已经被执行过了，直接return.
		return
	}
	// 调用f, 因为channel中只有一个值
	// 所以只有一个goroutine会到达这里
	f()

	// 关闭通道，这将释放所有在等待的
	// 以及未来会调用Do方法的goroutine
	close(o)
}

// 信号量（英语：semaphore）又称为信号标，是一个同步对象，用于保持在0至指定最大值之间的一个计数值。
// 当线程完成一次对该semaphore对象的等待（wait）时，该计数值减1；当线程完成一次对semaphore对象的释放（release）时，计数值加1。当计数值为0，则线程直至该semaphore对象变成signaled状态才能等待成功。semaphore对象的计数值大于0，为signaled状态；计数值等于0，为nonsignaled状态.
type Semaphore chan struct{}

func NewSemaphore(size int) Semaphore {
	return make(Semaphore, size)
}

func (s Semaphore) Lock() {
	// 只有在s还有空间的时候才能发送成功
	s <- struct{}{}
}

func (s Semaphore) Unlock() {
	// 为其他信号量腾出空间
	<-s
}

// 互斥锁是大小为1的信号量的特例。那么在刚才实现的信号量的基础上实现互斥锁只需要
type Mutex Semaphore

func NewMutex() Mutex {
	return Mutex(NewSemaphore(1))
}

// RWMutex是一个稍微复杂的原语：它允许任意数量的并发读锁，但在任何给定时间仅允许一个写锁,还可以保证，如果有线程持有写锁，则任何线程都不能持有或获得读锁。
// RWMutex具有三种状态：空闲，存在写锁和存在读锁。这意味着我们需要两个通道分别标记RWMutex上的读锁和写锁：空闲时，两个通道都为空；当获取到写锁时，标记写锁的通道里将被写入一下空结构体；当获取到读锁时，我们向两个通道中都写入一个值(避免写锁能够向标记写锁的通道发送值)，其中标记读锁的通道里的值代表当前RWMutex拥有的读锁的数量，读锁释放的时候除了更新通道里存的读锁数量值，也会抽空写锁通道。

type RWMutex struct {
	write   chan struct{}
	readers chan int
}

func NewLock() RWMutex {
	return RWMutex{
		// 用来做一个普通的互斥锁
		write: make(chan struct{}, 1),
		// 用来保护读锁的数量，获取读锁时通过接受通道里的值确保
		// 其他goroutine不会在同一时间更改读锁的数量。
		readers: make(chan int, 1),
	}
}

func (l RWMutex) Lock() {
	l.write <- struct{}{}
}

func (l RWMutex) Unlock() {
	<-l.write
}

func (l RWMutex) RLock() {
	// 统计当前读锁的数量，默认为0
	var rs int
	select {
	case l.write <- struct{}{}:
		// 如果write通道能发送成功，证明现在没有读锁
		// 向write通道发送一个值，防止出现并发的读-写
	case rs = <-l.readers:
		// 能从通道里接收到值，证明RWMutex上已经有读锁了，下面会更新读锁数量
	}
	// 如果执行了l.write <- struct{}{}, rs的值会是0
	rs++
	// 更新RWMutex读锁数量
	l.readers <- rs
}

func (l RWMutex) RUnlock() {
	// 读出读锁数量然后减一
	rs := <-l.readers
	rs--
	// 如果释放后读锁的数量变为0了，抽空write通道，让write通道变为可用
	if rs == 0 {
		<-l.write
		return
	}
	// 如果释放后读锁的数量减一后不是0，把新的读锁数量发送给readers通道
	l.readers <- rs
}

// WaitGroup最常见的用途是创建一个组，向其计数器中设置一个计数，生成与该计数一样多的goroutine，然后等待它们完成。每次goroutine运行完毕后，它将在组上调用Done表示已完成工作。可以通过调用WaitGroup的Done方法或以负数调用Add方法减少计数器的计数。当计数器达到0时，被Wait方法阻塞住的主线程会恢复执行。
//WaitGroup一个鲜为人知的功能是在计数器达到0后，如果调用Add方法让计数器变为正数，这将使WaitGroup重回阻塞状态。这意味着对于每个给定的WaitGroup，都有一点"世代"的意味：
//		当计数器从0移到正数时开始"世代"。
//		当计数器重回到0时，WaitGroup的一个世代结束。
//		当一个世代结束时，被该世代的所阻塞住的线程将恢复执行。

type generation struct {
	// 用于让等待者阻塞住的通道
	// 这个通道永远不会用于发送，只用于接收和close。
	wait chan struct{}
	// 计数器，标记需要等待执行完成的job数量
	n int
}

func newGeneration() generation {
	return generation{wait: make(chan struct{})}
}

func (g generation) end() {
	// close通道将释放因为接受通道而阻塞住的goroutine
	close(g.wait)
}

//这里我们使用一个通道来保护当前的generation。
//它基本上是WaitGroup状态的互斥量。
type WaitGroup chan generation

func NewWaitGroup() WaitGroup {
	wg := make(WaitGroup, 1)
	g := newGeneration()
	// 在一个新的WaitGroup上Wait, 因为计数器是0，会立即返回不会阻塞住线程
	// 它表现跟当前世代已经结束了一样, 所以这里先把世代里的wait通道close掉
	// 防止刚创建WaitGroup时调用Wait函数会阻塞线程
	g.end()
	wg <- g
	return wg
}

func (wg WaitGroup) Add(delta int) {
	//获取当前的世代
	g := <-wg
	if g.n == 0 {
		// 计数器是0，创建一个新的世代
		g = newGeneration()
	}
	g.n += delta
	if g.n <= 0 {
		// 跟sync库里的WaitGroup一样，不允许计数器为负数
		panic("negative WaitGroup count")
	}
	if g.n == 0 {
		// 计数器回到0了，关闭wait通道，被WaitGroup的Wait方法
		// 阻塞住的线程会被释放出来继续往下执行
		g.end()
	}

	// 将更新后的世代发送回WaitGroup通道
	wg <- g
}

func (wg WaitGroup) Done() {
	wg.Add(-1)
}

func (wg WaitGroup) Wait() {
	// 获取当前的世代
	g := <-wg
	// 保存一个世代里wait通道的引用
	wait := g.wait
	// 将世代写回WaitGroup通道
	wg <- g
	// 接收世代里的wait通道
	// 因为wait通道里没有值，会把调用Wait方法的goroutine阻塞住
	// 直到WaitGroup的计数器回到0，wait通道被close后才会解除阻塞
	<-wait
}
