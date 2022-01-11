package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const BindUrl = "https://login.uniontech.com"
const envPath = "/etc/environment"
const envDelimiter = "="

func GetServerAPI(env, value string) string {
	api := os.Getenv(env)
	if len(api) != 0 {
		return api
	}
	conf := readEtcEnvConf(env)
	if len(conf) != 0 {
		return conf
	}
	return value
}

// 兼容特殊场景下/etc/environment无法读取的问题
func readEtcEnvConf(env string) string {
	fi, err := os.Open(envPath)
	if err != nil {
		return ""
	}
	defer fi.Close()
	line := bufio.NewReader(fi)
	for {
		content, _, err := line.ReadLine()
		if err == io.EOF {
			break
		}
		if content == nil {
			continue
		}
		contentStr := string(content)
		if strings.HasPrefix(contentStr, env) {
			path := strings.Split(contentStr, envDelimiter)
			return strings.TrimSpace(path[1])
		}
	}
	return ""
}

func main() {
	s := GetServerAPI("DEEPINID_OAUTH_URI", BindUrl) + "/api/key"
	fmt.Println(s)
}
