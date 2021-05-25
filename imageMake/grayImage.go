package main

import (
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func main() {
	imagePath := "origin.png"
	fi, err := os.Open(imagePath)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	decode, _, err := image.Decode(fi)
	if err != nil {
		panic(err)
	}
	rgba := grayingImage(decode)
	newFileName := "gray_png.png"
	outFile, err := os.Create(newFileName)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	if err := imageEncode(newFileName, outFile, rgba); err != nil {
		panic(err)
	}
}

func grayingImage(m image.Image) *image.RGBA {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	newRgba := image.NewRGBA(bounds)
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			colorRgb := m.At(x, y)
			_, g, _, a := colorRgb.RGBA()
			newG := uint8(g >> 8)
			newA := uint8(a >> 8)
			// 将每个点的设置为灰度值
			newRgba.SetRGBA(x, y, color.RGBA{R: newG, G: newG, B: newG, A: newA})
		}
	}
	return newRgba
}

func imageEncode(fileName string, file *os.File, rgba *image.RGBA) error {
	// 将图片和扩展名分离
	stringSlice := strings.Split(fileName, ".")
	// 根据图片的扩展名来运用不同的处理
	switch stringSlice[len(stringSlice)-1] {
	case "jpg":
		return jpeg.Encode(file, rgba, nil)
	case "jpeg":
		return jpeg.Encode(file, rgba, nil)
	case "gif":
		return gif.Encode(file, rgba, nil)
	case "png":
		return png.Encode(file, rgba)
	default:
		panic("不支持的图片类型")
	}
}
