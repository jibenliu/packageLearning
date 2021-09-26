package main

import (
	"context"
	"fmt"
	"golang.org/x/sys/unix"
	"net"
	"net/http"
	"os"
	"syscall"
)

var lc = net.ListenConfig{
	Control: func(network, address string, c syscall.RawConn) error {
		var opErr error
		if err := c.Control(func(fd uintptr) {
			opErr = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1)
		}); err != nil {
			return err
		}
		return opErr
	},
}

func main() {
	pid := os.Getpid()
	l, err := lc.Listen(context.Background(), "tcp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}
	server := &http.Server{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Client [%s] Received msg from Server PID: [%d] \n", r.RemoteAddr, pid)
	})
	fmt.Printf("Server with PID: [%d] is running \n", pid)
	_ = server.Serve(l)
}
