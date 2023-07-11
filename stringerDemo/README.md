# 安装
```shell
go install golang.org/x/tools/cmd/stringer@latest
```
#检查
```shell
stringer -h
```

# 编写
```go
// code.go
package mycodes

type ErrCode int

const (
	SUCCESS        ErrCode = 200
	BadRequest     ErrCode = 400
	Unauthorized   ErrCode = 401
	Forbidden      ErrCode = 403
	TooManyRequest ErrCode = 429
	InternalServer ErrCode = 500
)

const (
	// 这里填写业务错误码，从 10000 开始
	UsernameOrPassword                 ErrCode = iota + 10001 // 用户名或密码不正确
	UserClosed                                                // 用户被禁用了
	NoRegister                                                // 用户还没有注册
	PhoneExist                                                // 手机号已经存在
	VerifyCodeExpired                                         // 验证码超过有效期
	ErrorVerifyCode                                           // 验证码错误
	VerifyCodeHasUsed                                         // 验证码已使用
	ErrorSaveAvatar                                           // 保存头像出错
	ErrorConversationReachedTrialLimit                        // 达到会话数量限制

	ErrorPlaceOrder     // 订单创建失败
	ErrorNotifyOrder    // 订单创建失败
	ErrorContentIllegal // 内容非法

	ErrorGPT4ConversationReachedDailyLimit // 达到gpt4会话每日数量限制
	ErrorNotVipGPT4ConversationNotAllowed  // 非会员禁止发起gpt4会话

	ErrorCreditReachedLimit // 剩余字符不足
	ErrorNotAllowedEngine   // 非vip用户禁止试用引擎
)

// main.go
package main

import (
    "fmt"
    "stringerDemo/mycodes"
)

//go:generate stringer -linecomment -type ErrCode ./mycodes
func main() {
	fmt.Println(mycodes.BadRequest)
}
```

# 运行
```shell
go generate -x
```
