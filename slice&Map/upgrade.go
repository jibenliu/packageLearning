package main

import (
	"fmt"
	"strings"
)

type Server struct {
	Name string
}

type Servers []Server

// ListServers 返回服务器列表
func ListServers() Servers {
	return []Server{
		{Name: "Server1"},
		{Name: "Server2"},
		{Name: "Foo1"},
		{Name: "Foo2"},
	}
}

// Filter 返回包含 name 的服务器。空的 name 将会返回所有服务器。
func (s Servers) filterServers(name string) Servers {
	filtered := make(Servers, 0)

	for _, server := range s {
		if strings.Contains(server.Name, name) {
			filtered = append(filtered, server)
		}

	}

	return filtered
}

func main() {
	servers := ListServers()
	servers = servers.filterServers("Foo")
	fmt.Printf("servers %+vn", servers)
}