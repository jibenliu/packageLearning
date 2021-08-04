package main

import (
	"errors"
	"fmt"
	"go.uber.org/dig"
	"net/http"
	"strconv"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type ServerConfig struct {
	Host        string  // 主机地址
	Port        string  // 端口号
	Used        bool    // 是否被占用
	mongoConfig *Config `optional:"true"`
}

type ServerGroup struct {
	dig.In

	Servers []*Server `group:"server"`
}

type ServerConfigNamed struct {
	dig.In

	Config1 *ServerConfig `name:"config1"`
	Config2 *ServerConfig `name:"config2"`
	Config3 *ServerConfig `name:"config3"`
}

type Server struct {
	Config *ServerConfig
}

func (s *Server) Run(i int) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(write http.ResponseWriter, req *http.Request) {
		write.Write([]byte(fmt.Sprintf("第%s个服务，端口号是: %s", strconv.FormatInt(int64(i), 10), s.Config.Port)))
	})
	http.ListenAndServe(s.Config.Host+":"+s.Config.Port, mux)
}

func NewServerConfig(port string) func() *ServerConfig {
	return func() *ServerConfig {
		return &ServerConfig{Host: "127.0.0.1", Port: port, Used: false}
	}
}

func NewServer(sc ServerConfigNamed) *Server {
	if !sc.Config1.Used {
		sc.Config1.Used = true
		return &Server{Config: sc.Config1}
	} else if !sc.Config2.Used {
		sc.Config2.Used = true
		return &Server{Config: sc.Config2}
	} else if !sc.Config3.Used {
		sc.Config3.Used = true
		return &Server{Config: sc.Config3}
	}
	panic(errors.New("none config is captured"))
}

func ServerRun(sg ServerGroup) {
	for i, s := range sg.Servers {
		go s.Run(i)
	}
}

func main() {
	container := dig.New()
	// 注入 3 个服务配置项
	container.Provide(NewServerConfig("8199"), dig.Name("config1"))
	container.Provide(NewServerConfig("8299"), dig.Name("config2"))
	container.Provide(NewServerConfig("8399"), dig.Name("config3"))
	// 注入 3 个服务实例
	container.Provide(NewServer, dig.Group("server"))
	container.Provide(NewServer, dig.Group("server"))
	container.Provide(NewServer, dig.Group("server"))
	// 使用缓冲 channel 卡住主协程
	serverChan := make(chan int, 1)
	container.Invoke(ServerRun)
	<-serverChan
}
