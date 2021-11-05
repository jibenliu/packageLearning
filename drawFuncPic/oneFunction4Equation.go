package main

import "github.com/fogleman/gg"

func main() {
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SavePNG("./images/oneFunctionEquation.png")
}
