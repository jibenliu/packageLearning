package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
)

func getMacAddress() (macAddress []string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("fail to get net interfaces:%v", err)
		return macAddress
	}
	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}
		macAddress = append(macAddress, macAddr)
	}
	return macAddress
}

func getIps() (ips []string) {
	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("fail to get net interface address:%v", err)
		return ips
	}
	for _, address := range interfaceAddr {
		ipNet, isValidIpNet := address.(*net.IPNet)
		if isValidIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return ips
}

func getLocalIPv4() (syscall.SockaddrInet4, error) {
	// 创建一个socket连接对象
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, 0)
	if err != nil {
		return syscall.SockaddrInet4{}, err
	}
	defer syscall.Close(fd)

	// 连接到一个私网地址
	err = syscall.Connect(fd, &syscall.SockaddrInet4{
		Addr: [4]byte{10, 255, 255, 255},
	})
	if err != nil {
		return syscall.SockaddrInet4{}, err
	}
	// 获取连接本地信息
	addr, err := syscall.Getsockname(fd)
	if err != nil {
		return syscall.SockaddrInet4{}, err
	}
	return *addr.(*syscall.SockaddrInet4), nil
}

// GetProcessName 获取进程名称
func GetProcessName() string {
	processName := os.Args[0]
	//fmt.Printf("processName: %v\n", processName)
	// C:\Users\admin\Desktop\cmd.exe
	// ./cmd
	slashPos := strings.LastIndex(processName, "\\")
	if slashPos > 0 {
		return processName[slashPos+1:]
	}
	backslashPos := strings.LastIndex(processName, "/")
	if backslashPos > 0 {
		return processName[backslashPos+1:]
	}
	return "unknown"
}

// GetAppPath
// when the project is running on build way to get path
func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}

// GetComputerName 获取机器名
func GetComputerName() string {
	sHostName, _ := os.Hostname()
	fmt.Printf("sHostName: %v\n", sHostName)
	// message too long for RSA public key size
	if len(sHostName) > 10 {
		sHostName = sHostName[1 : 10-1]
	}
	if runtime.GOOS == "linux" {
		sHostName = sHostName + " (Linux)"
	} else if runtime.GOOS == "darwin" {
		sHostName = sHostName + " (Darwin)"
	}
	return sHostName
}

func main() {
	fmt.Printf("mac address:%q\n", getMacAddress())
	fmt.Printf("ip address:%q\n", getIps())
	ipv4, err := getLocalIPv4()
	fmt.Printf("ip address:%q err is %s\n", ipv4, err)
	fmt.Printf("process name:%q\n", GetProcessName())
	fmt.Printf("computer name:%q\n", GetComputerName())
	fmt.Printf("app path:%q\n", GetAppPath())
}
