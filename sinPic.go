package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

const size = 300

func main() {
	pic := image.NewGray(image.Rect(0, 0, size, size))

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			pic.SetGray(x, y, color.Gray{Y: 255})
		}
	}

	for x := 0; x < size; x++ {
		s := float64(x) * 2 * math.Pi / size

		y := size/2 - math.Sin(s)*size/2

		pic.SetGray(x, int(y), color.Gray{Y: 0})
	}

	file, err := os.Create("sin.png")
	if err != nil {
		log.Fatal(err)
	}
	nErr := png.Encode(file, pic)
	if nErr != nil{
		log.Fatal(nErr)
	}

	file.Close()
}
