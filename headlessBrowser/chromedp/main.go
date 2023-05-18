package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func main() {
	// 创建一个上下文
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// 设置浏览器选项
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("remote-debugging-port", "9222"),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(ctx, opts...)
	defer cancel()

	// 创建一个浏览器实例
	ctx, cancel = chromedp.NewContext(allocCtx)
	defer cancel()

	// 导航到指定的URL
	var buf []byte
	err := chromedp.Run(ctx, chromedp.Navigate("https://www.baidu.com"), chromedp.Sleep(2*time.Second), chromedp.ActionFunc(func(ctx context.Context) error {
		// 获取页面截图
		var err error
		buf, err = page.CaptureScreenshot().WithQuality(90).WithClip(&page.Viewport{X: 0, Y: 0, Width: 1920, Height: 1080, Scale: 1}).Do(ctx)
		if err != nil {
			return err
		}
		return nil
	}))
	if err != nil {
		log.Fatal(err)
	}
	// 将截图保存到文件
	err = os.WriteFile("screenshot.png", buf, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
