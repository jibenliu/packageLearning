package main

import (
	"auth_socks5/socks5"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	fmt.Println("默认端口：9091，无账号密码")
	port := os.Getenv("SOCKS5_PORT")
	if port == "" {
		port = "9091"
	}
	conf := &socks5.Config{
		Logger: log.New(os.Stdout, "", log.Llongfile),
	}
	username := os.Getenv("SOCKS5_USERNAME")
	password := os.Getenv("SOCKS5_PASSWORD")
	if username != "" {
		creds := socks5.StaticCredentials{
			username: password,
		}
		conf.Credentials = creds
	}
	fmt.Printf("当前端口：%s，账号: %s 密码: %s\n", port, username, password)
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}
	if err := server.ListenAndServe("tcp", ":"+port); err != nil {
		panic(err)
	}
}
