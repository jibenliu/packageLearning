package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"syscall"
	"time"
)

type gracefulListener struct {
	net.Listener
	stop    chan error
	stopped bool
}

var httpWg sync.WaitGroup

func (gl *gracefulListener) Accept() (c net.Conn, err error) {
	c, err = gl.Listener.Accept()
	if err != nil {
		return
	}

	c = gracefulConn{Conn: c}

	httpWg.Add(1)
	return
}

func (gl *gracefulListener) Close() error {
	if gl.stopped {
		return syscall.EINVAL
	}
	gl.stop <- nil
	return <-gl.stop
}

func (gl *gracefulListener) File() *os.File {
	tl := gl.Listener.(*net.TCPListener)
	fl, _ := tl.File()
	return fl
}

type gracefulConn struct {
	net.Conn
}

func (w gracefulConn) Close() error {
	httpWg.Done()
	return w.Conn.Close()
}

func newGracefulListener(l net.Listener) (gl *gracefulListener) {
	gl = &gracefulListener{Listener: l, stop: make(chan error)}
	go func() {
		_ = <-gl.stop
		gl.stopped = true
		gl.stop <- gl.Listener.Close()
	}()
	return
}

func main() {
	server := &http.Server{
		Addr:           "127.0.0.1:8888",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 16,
	}

	var gracefulChild bool
	var l net.Listener
	var err error

	flag.BoolVar(&gracefulChild, "graceful", false, "listen on fd open 3 (internal use only)")

	if gracefulChild {
		log.Print("main: Listening to existing file descriptor 3.")
		f := os.NewFile(3, "")
		l, err = net.FileListener(f)
	} else {
		log.Print("main: Listening on a new file descriptor.")
		l, err = net.Listen("tcp", server.Addr)
	}
	if err != nil {
		log.Fatalf("error is%v", err)
	}

	if gracefulChild {
		parent := syscall.Getppid()
		log.Printf("main: Killing parent pid: %v", parent)
		syscall.Kill(parent, syscall.SIGTERM)
	}

	netListener := newGracefulListener(l)
	server.Serve(netListener)
}
