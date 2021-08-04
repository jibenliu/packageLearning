package main

import (
	"fmt"
	"net"
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

func main() {
	fmt.Printf("mac address:%q\n", getMacAddress())
	fmt.Printf("ip address:%q\n", getIps())
}
