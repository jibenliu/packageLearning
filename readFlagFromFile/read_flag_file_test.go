package readFlagFromFile

import (
	"flag"
	"strconv"
	"testing"
)

var (
	RedisAddress string
	RedisPort    int
	TestQuotaStr string
	TestStr      string
)

func init() {
	flag.StringVar(&RedisAddress, "redis_address", "127.0.0.1", "this is redis address")
	flag.IntVar(&RedisPort, "redis_port", 6379, "this is redis port")
	flag.StringVar(&TestQuotaStr, "aaa", "abdc", "this is quota str")
	flag.StringVar(&TestStr, "bbb", "safd", "this is undefined str")
}

// 测试传入参数 -flagfile=conf.ini
func TestParse(t *testing.T) {
	flagFile = "conf.ini"
	if err := Parse(); err != nil {
		t.Errorf("Parse() error = %v, wantErr %v", err, nil)
	}
	t.Logf(RedisAddress)
	t.Logf(strconv.Itoa(RedisPort))
	t.Logf(TestQuotaStr)
	t.Logf(TestStr)
}
