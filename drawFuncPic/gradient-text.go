package main

import (
	"image/color"

	"github.com/fogleman/gg"
)

const (
	W = 1024
	H = 512
)

func main() {
	dc := gg.NewContext(W, H)

	// draw text
	dc.SetRGB(0, 0, 0)
	dc.LoadFontFace("./ttf/ray.ttf", 128)
	dc.DrawStringAnchored("Gradient Text", W/2, H/2, 0.5, 0.5)

	// get the context as an alpha mask
	mask := dc.AsMask()

	// clear the context
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// set a gradient
	g := gg.NewLinearGradient(0, 0, W, H)
	g.AddColorStop(0, color.RGBA{R: 255, A: 255})
	g.AddColorStop(1, color.RGBA{B: 255, A: 255})
	dc.SetFillStyle(g)

	// using the mask, fill the context with the gradient
	dc.SetMask(mask)
	dc.DrawRectangle(0, 0, W, H)
	dc.Fill()

	dc.SavePNG("./images/gradient-text.png")
}
