package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config 定义config结构体
type Config struct {
	AppId  string
	Secret string
	Host   Host
}

// Host json中的嵌套对应结构体的嵌套
type Host struct {
	Address string
	Port    int
}

func main() {
	config := viper.New()
	config.WatchConfig()
	config.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	config.AddConfigPath("./")
	config.SetConfigName("config")
	config.SetConfigType("json")
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	fmt.Println(config.GetString("appId"))
	fmt.Println(config.GetString("secret"))
	fmt.Println(config.GetString("host.address"))
	fmt.Println(config.GetString("host.port"))

	//直接反序列化为Struct
	var configJson Config
	if err := config.Unmarshal(&configJson); err != nil {
		fmt.Println(err)
	}

	fmt.Println(configJson.Host)
	fmt.Println(configJson.AppId)
	fmt.Println(configJson.Secret)

	_ = config.SafeWriteConfigAs("config-bak.json")

}
