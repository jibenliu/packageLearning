package util

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type HttpConPool struct {
	Conn *http.Client
}

var Hpool *HttpConPool

func loadHttpPool() {
	Hpool = new(HttpConPool)
	Hpool.Conn = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: Str2int(GetConfigValue("http", "max_conn")),
		},
		Timeout: time.Duration(Str2int(GetConfigValue("http", "timeout"))) * time.Millisecond,
	}
}
func GetConfigValue(name string, expiredKey string) string {
	return ""
}
func Str2int(name string) int {
	num, _ := strconv.Atoi(name)
	return num
}
func (h *HttpConPool) Request(url string, method string, data string, header map[string]string) (interface{}, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return nil, err
	}

	for h, v := range header {
		req.Header.Set(h, v)
	}

	response, err := h.Conn.Do(req)

	if err != nil {
		return nil, err
	} else if response != nil {
		defer response.Body.Close()

		rBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		} else {
			return string(rBody), nil
		}
	} else {
		return nil, nil
	}
}
