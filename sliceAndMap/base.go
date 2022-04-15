package main

import (
	"fmt"
	"strings"
)

type Server struct {
	Name string
}

func ListServers() []Server {
	return []Server{
		{Name: "Server1"},
		{Name: "Server2"},
		{Name: "Foo1"},
		{Name: "Foo2"},
	}
}

func filterServers(name string) []Server {
	servers := ListServers()

	// 返回所有服务器
	if name == "" {
		return servers
	}

	// 返回过滤后的结果
	filtered := make([]Server, 0)

	for _, server := range servers {
		if strings.Contains(server.Name, name) {
			filtered = append(filtered, server)
		}
	}

	return filtered
}

func main() {
	servers := filterServers("Foo")

	// 输出: "servers [{Name:Foo1} {Name:Foo2}]"
	fmt.Printf("servers %+vn", servers)
}
