package main

import (
	"flag"
	"fmt"
	"github.com/FZambia/sentinel"
	"github.com/gomodule/redigo/redis"

	"strings"
	"time"
)

var RedisConnPool *redis.Pool

func init() {
	redisAddr := flag.String("redisAddr", "", "redis addr")
	redisType := flag.String("redisType", "", "redis type cluster or sentinel")
	MaxIdle := flag.Int("maxIdle", 0, "max idle connection")
	masterName := flag.String("masterName", "master-redis", "master node name")
	password := flag.String("redisPassword", "", "password")
	flag.Parse()
	if *redisType == "sentinel" {
		redisAddrs := strings.Split(*redisAddr, ",")
		sntnl := &sentinel.Sentinel{
			Addrs:      redisAddrs,
			MasterName: *masterName,
			Dial: func(addr string) (redis.Conn, error) {
				timeout := 500 * time.Millisecond
				c, err := redis.DialTimeout("tcp", addr, timeout, timeout, timeout)
				if err != nil {
					return nil, err
				}
				return c, nil
			},
		}

		RedisConnPool = &redis.Pool{
			MaxIdle:     *MaxIdle,
			IdleTimeout: 240 * time.Second,
			Dial: func() (redis.Conn, error) {
				masterAddr, err := sntnl.MasterAddr()
				if err != nil {
					return nil, err
				}
				c, err := redis.Dial("tcp", masterAddr)
				if err != nil {
					return nil, err
				}
				//此处1234对应redis密码
				if _, err := c.Do("AUTH", *password); err != nil {
					c.Close()
					return nil, err
				}
				return c, nil
			},
			TestOnBorrow: CheckRedisRole,
		}
	} else {
		RedisConnPool = &redis.Pool{
			MaxIdle:     *MaxIdle,
			IdleTimeout: 240 * time.Second,
			// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", *redisAddr)
				if err != nil {
					c.Close()
					return nil, err
				}
				//此处1234对应redis密码
				if _, err := c.Do("AUTH", *password); err != nil {
					c.Close()
					return nil, err
				}
				return c, nil
			},
		}
	}
}

func CheckRedisRole(c redis.Conn, t time.Time) error {
	if !sentinel.TestRole(c, "master") {
		return fmt.Errorf("role check failed")
	} else {
		return nil
	}
}
func main() {
	conn := RedisConnPool.Get()
	defer conn.Close()
	_, _ = redis.String(conn.Do("GET", "aaa"))

}
