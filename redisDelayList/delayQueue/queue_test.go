package delayQueue

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_DelayJob(t *testing.T) {
	ASSERT := assert.New(t)
	InitRedis()
	q := queue{}
	job := PrintJob{}
	q.DelayJob(&Job{
		Id:         "1",
		Topic:      job.Topic(),
		Delay:      time.Now().Unix() + 50,
		PlayLoad:   []byte("处理延时任务!"),
		Timestamp:  time.Now().Unix(),
		TryCount:   0,
		TryTimeOut: 10,
	})
	_, ct, _ := q.GetJobNum(job.Topic())
	ASSERT.Equal(1, ct, "ERRROR NUM")
	time.Sleep(time.Second * 52)
	_, ct, _ = q.GetJobNum(job.Topic())
	ASSERT.Equal(0, ct, "ERRROR NUM")
}
