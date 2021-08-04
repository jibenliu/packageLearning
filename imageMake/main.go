package main

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
)

// Jpeg的图片压缩是很好做的，因为jpeg这个协议本身就支持调整图片质量的。在golang中，我们只需要使用标准库的image/jpeg，将图片从二进制数据解码后，降低质量再编码为二进制数据即可实现压缩。而且质量和压缩比例相对而言还不错。
// 比较麻烦的是压缩PNG图片，后来，借鉴一篇博客的做法，还是先把PNG图片转换为Jpeg图片，然后再将jpeg图片的质量降低
func compressImagesResource(data []byte) []byte {
	imgSrc, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		panic(err)
		return data
	}
	newImg := image.NewRGBA(imgSrc.Bounds())
	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{C: color.White}, image.Point{}, draw.Src)
	draw.Draw(newImg, newImg.Bounds(), imgSrc, imgSrc.Bounds().Min, draw.Over)
	buf := bytes.Buffer{}
	err = jpeg.Encode(&buf, newImg, &jpeg.Options{Quality: 40})
	if err != nil {
		panic(err)
		return data
	}
	if buf.Len() > len(data) {
		panic("ssss")
		return data
	}
	return buf.Bytes()
}

func main() {
	data, err := ioutil.ReadFile("origin.png")
	if err != nil {
		panic(err)
	}
	bts := compressImagesResource(data)
	err = ioutil.WriteFile("bak.jpg", bts, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
