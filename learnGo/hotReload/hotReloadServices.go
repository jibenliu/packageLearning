package hotReload

import (
	"flag"
	"net"
	"net/http"
	"time"
)

/**
* 处理过程分为以下几个步骤：
* 	监听信号（USR2..）
* 	收到信号时fork子进程（使用相同的启动命令），将服务监听的socket文件描述符传递给子进程
* 	子进程监听父进程的socket，这个时候父进程和子进程都可以接收请求
* 	子进程启动成功之后，父进程停止接收新的连接，等待旧连接处理完成（或超时）
*	父进程退出，重启完成
* 细节
* 	父进程将socket文件描述符传递给子进程可以通过命令行，或者环境变量等
* 	子进程启动时使用和父进程一样的命令行，对于golang来说用更新的可执行程序覆盖旧程序
* 	server.Shutdown()优雅关闭方法是go>=1.8的新特性
* 	server.Serve(l)方法在Shutdown时立即返回，Shutdown方法则阻塞至context完成，所以Shutdown的方法要写在主goroutine中
 */

var (
	server   *http.Server
	listener net.Listener
	graceful = flag.Bool("graceful", false, "listen on fd open 3 (internal use only)")
)

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(20 * time.Second)
	_, _ = w.Write([]byte("hello world!"))
}

func main()  {
	flag.Parse()
}

