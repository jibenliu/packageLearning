package delayQueue

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"time"
)

var redisDB *redis.Client

func SetRedis(db *redis.Client) {
	redisDB = db
}

func InitRedis() {
	redisDB = redis.NewClient(&redis.Options{
		Addr:         "192.168.1.51:32353",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
	redisDB.WrapProcess(func(old func(cmd redis.Cmder) error) func(cmd redis.Cmder) error {
		return func(cmd redis.Cmder) error {
			log.Debugf("starting processing: <%s>\n", cmd)
			err := old(cmd)
			log.Debugf("finished processing: <%s>\n", cmd)
			return err
		}
	})

	_, err := redisDB.Ping().Result()
	if err != nil {
		panic(err)
	}
}

// zcard 获取延时消息数
func zcard(key string) *redis.IntCmd {
	return redisDB.ZCard(key)
}

// list 获取等待执行的消息数
func llen(key string) *redis.IntCmd {
	return redisDB.LLen(key)
}

// lpush  推送新的job到队列
func lpush(key string, value interface{}) error {
	return redisDB.LPush(key, value).Err()
}

// zadd 增加延迟job
func zadd(key string, value interface{}, delay int) error {
	return redisDB.ZAdd(key, redis.Z{Score: float64(delay), Member: value}).Err()
}

// rpop 从ready队列取消息
func rpop(key string) *redis.StringCmd {
	return redisDB.RPop(key)
}

// 将到期的job迁移到ready队列等待执行，这里使用redis script实现
func migrateExpiredJobs(delaykey, readyKey string) error {
	script := redis.NewScript(`
  local val = redis.call('zrangebyscore', KEYS[1], '-inf', ARGV[1], 'limit', 0, 20)
  if(next(val) ~= nil) then
    redis.call('zremrangebyrank', KEYS[1], 0, #val - 1)
    redis.call('rpush', KEYS[2], unpack(val, 1, #val))
  end
  return #val
  `)
	return script.Run(redisDB, []string{delaykey, readyKey}, time.Now().Unix()).Err()
}
