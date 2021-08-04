package collectGroupV2

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestWithContext(t *testing.T) {

	// 创建一个errGroup
	group, ctx := WithContext(context.Background())
	// 模拟多任务
	tasks := []task{
		func() error {
			time.Sleep(4 * time.Second)
			t.Log("向订单表加入消息....")
			return nil
		},
		func() error {
			time.Sleep(2 * time.Second)
			t.Log("更新库存消息....")
			return nil
		},
		func() error {
			time.Sleep(3 * time.Second)
			t.Log("发送用户通知.....")
			return nil
		},
		func() error {
			time.Sleep(4 * time.Second)
			return errors.New("用户扣款发送错误")
		},
	}

	for i, t := range tasks {
		group.Go(fmt.Sprintf("go-id-%s", strconv.Itoa(i)), t)
	}
	// group.Wait()
	// 监听任务出错了一个就返回
	<-ctx.Done()
	if len(group.Errs) > 0 {
		t.Log("group exit...任务出，拿到错误消息回滚业务....")
		t.Log("Get errors: ", group.Errs)
	}
}
