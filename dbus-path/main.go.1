package main

import (
	"flag"
	"fmt"
	"github.com/godbus/dbus/v5"
	"github.com/jandre/procfs"
	"os"
)

var (
	senderId = flag.String("senderId", "", "the sender id")
)

func main() {
	flag.Parse()
	fmt.Println(*senderId)
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to session bus:", err)
		os.Exit(1)
	}
	defer conn.Close()

	var s []string
	err = conn.BusObject().Call("org.freedesktop.DBus.ListNames", 0).Store(&s)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to get list of owned names:", err)
		os.Exit(1)
	}

	for _, v := range s {
		if v == *senderId {
			var id int
			err = conn.BusObject().Call("org.freedesktop.DBus.GetConnectionUnixProcessID", 0, v).Store(&id)
			path, _ := GetProcessPath(id)
			fmt.Printf("sender id is: %s, ------ path is %s:\n", v, path)
		}
	}
}

func GetProcessPath(pid int) (exe string, err error) {
	process, err := procfs.NewProcess(pid, true)
	if err != nil {
		return "", err
	}
	return process.Exe, nil
}
