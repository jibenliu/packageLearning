package main

import (
	"github.com/dchest/captcha"
	"github.com/go-redis/redis"
	"os"
)

type redisMem struct {
	Client *redis.Client
	Extra  interface{}
}

var redisDb *redisMem

// 初始化连接
func initClient() (err error) {
	redisDb = &redisMem{
		Client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0, // use default DB
		}),
	}

	_, err = redisDb.Client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
func (r *redisMem) Set(id string, digits []byte) {
	r.Client.Set(id, string(digits), captcha.Expiration)
}
func (r *redisMem) Get(id string, clear bool) (digits []byte) {
	res, _ := r.Client.Get(id).Result()
	digits = []byte(res)
	return
}

func main() {
	//err := initClient()
	//if err != nil{
	//	fmt.Println(err)
	//	os.Exit(0)
	//}
	for i := 0; i < 100; i++ {
		//captcha.SetCustomStore(redisDb)
		id := captcha.New()
		//captchaStr := redisDb.Get(id, false)
		//fmt.Println(captcha.VerifyString(id, string(captchaStr)))

		f, _ := os.Create("./images/" + id + ".png") //创建文件
		//f, _ := os.Create("./images/test.png") //创建文件
		defer f.Close()
		_ = captcha.WriteImage(f, id, captcha.StdWidth, captcha.StdHeight)
	}
}
