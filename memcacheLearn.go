package main

import (
	"fmt"
	"sync"
	"time"
)

type memoryCache struct {
	lock  *sync.RWMutex
	items map[interface{}]interface{}
}

func (mc *memoryCache) set(key interface{}, value interface{}) error {
	mc.lock.Lock()
	defer mc.lock.Unlock()
	mc.items[key] = value
	return nil
}

func (mc *memoryCache) get(key interface{}) interface{} {
	mc.lock.RLock()
	defer mc.lock.RUnlock()

	if val, ok := mc.items[key]; ok {
		return val
	}
	return nil
}

var gAccId2Token = &memoryCache{
	lock:  new(sync.RWMutex),
	items: make(map[interface{}]interface{}),
}

func getTokenByHttpUrl(url string) (interface{}, error) {
	//此处设计从http获取token的逻辑
	var token interface{}
	return token, nil
}
func GetTokenByAccountId(accountId uint) (map[string]interface{}, error) {
	token := gAccId2Token.get(accountId)
	if token != nil {
		now := time.Now().Unix()
		if int(now) < int(token.(map[string]interface{})["expireData"].(float64)) {
			return token.(map[string]interface{}), nil
		}
	}
	var apiUrl string
	token, err := getTokenByHttpUrl(apiUrl)
	if err != nil {
		return map[string]interface{}{}, nil
	}
	_ = gAccId2Token.set(accountId, token)
	return token.(map[string]interface{}), nil
}

func main() {
	token, _ := GetTokenByAccountId(1)
	fmt.Println(token)
}
