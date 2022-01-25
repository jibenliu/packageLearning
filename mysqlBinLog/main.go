package main

import (
	"flag"
	"fmt"
	"mysqlBinLog/app"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var myHost = flag.String("host", "127.0.0.1", "MySQL replication host")
var myPort = flag.Int("port", 3306, "MySQL replication port")
var myUser = flag.String("user", "root", "MySQL replication user")
var myPass = flag.String("pass", "****", "MySQL replication pass")
var serverId = flag.Int("server_id", 1111, "MySQL replication server id")

func main() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		os.Kill,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGQUIT,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	runtime.GOMAXPROCS(runtime.NumCPU()/4 + 1)
	flag.Parse()
	cfg := &app.Config{
		Host:     *myHost,
		Port:     *myPort,
		User:     *myUser,
		Pass:     *myPass,
		ServerId: *serverId,
		LogFile:  "mysql-bin.000032",
		Position: 3070,
	}
	srv := &app.Server{Cfg: cfg}
	go srv.Run()

	select {
	case n := <-sc:
		srv.Quit()
		fmt.Printf("receive signal %v, closing", n)
	}
}
