package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/rod/lib/utils"
	"github.com/ysmood/gson"
)

func main() {
	page := rod.New().MustConnect().MustPage("https://baidu.com").MustWaitLoad()

	//简单模式，默认截图设置
	page.MustScreenshot("my.png")

	//自定义截图设置
	img, _ := page.Screenshot(true, &proto.PageCaptureScreenshot{
		Format:  proto.PageCaptureScreenshotFormatJpeg,
		Quality: gson.Int(90),
		Clip: &proto.PageViewport{
			X:      0,
			Y:      0,
			Width:  300,
			Height: 200,
			Scale:  1,
		},
		FromSurface: true,
	})
	_ = utils.OutputFile("my.jpg", img)
}
