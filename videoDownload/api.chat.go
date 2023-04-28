package main

import (
	"github.com/henrylee2cn/pholcus/app/downloader/request"
	"github.com/henrylee2cn/pholcus/app/spider"
	"net/http"
)

func init() {
	spider.Spider{
		Name:        "静态规则示例",
		Description: "静态规则示例 [Auto Page] [http://xxx.xxx.xxx]",
		// Pausetime: 300,
		// Limit:   LIMIT,
		// Keyin:   KEYIN,
		EnableCookie:    true,
		NotDefaultField: false,
		Namespace:       nil,
		SubNamespace:    nil,
		RuleTree: &spider.RuleTree{
			Root: func(ctx *spider.Context) {
				ctx.AddQueue(&request.Request{Url: "http://xxx.xxx.xxx", Rule: "登录页"})
			},
			Trunk: map[string]*spider.Rule{
				"登录页": {
					ParseFunc: func(ctx *spider.Context) {
						ctx.AddQueue(&request.Request{
							Url:      "http://xxx.xxx.xxx",
							Rule:     "登录后",
							Method:   "POST",
							PostData: "username=123456@qq.com&password=123456&login_btn=login_btn&submit=login_btn",
						})
					},
				},
				"登录后": {
					ParseFunc: func(ctx *spider.Context) {
						ctx.Output(map[string]interface{}{
							"全部": ctx.GetText(),
						})
						ctx.AddQueue(&request.Request{
							Url:    "http://accounts.xxx.xxx/member",
							Rule:   "个人中心",
							Header: http.Header{"Referer": []string{ctx.GetUrl()}},
						})
					},
				},
				"个人中心": {
					ParseFunc: func(ctx *spider.Context) {
						ctx.Output(map[string]interface{}{
							"全部": ctx.GetText(),
						})
					},
				},
			},
		},
	}.Register()
}
