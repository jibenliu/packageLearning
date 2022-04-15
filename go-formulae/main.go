package main

import (
	"github.com/tdewolff/formulae"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"log"
)

func main() {
	// Parse formula
	in := "sin(cos(x))^2+1/x-1"
	f, errs := formulae.Parse(in)
	if len(errs) > 0 {
		log.Fatal(errs)
	}
	df := f.Derivative()

	// Calculate function
	xs, ys, errs := f.Interval(0.5, 0.01, 5.0)
	if len(errs) > 0 {
		log.Fatal(errs)
	}
	xys := make(plotter.XYs, len(xs))
	for i := range xs {
		xys[i].X = xs[i]
		xys[i].Y = real(ys[i])
	}
	_, _, ymin, _ := plotter.XYRange(xys)
	if ymin > 0 {
		ymin = 0
	}

	// Calculate function derivative
	xs2, ys2, errs := df.Interval(0.5, 0.01, 5.0) //计算导数
	if len(errs) > 0 {
		log.Fatal(errs)
	}
	xys2 := make(plotter.XYs, len(xs2))
	for i := range xs2 {
		xys2[i].X = xs2[i]
		xys2[i].Y = real(ys2[i])
	}

	// Plot functions
	p := plot.New()

	p.Title.Text = "Formula"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"
	p.Y.Min = ymin

	line, err := plotter.NewLine(xys)
	if err != nil {
		log.Fatal(err)
	}
	line2, err := plotter.NewLine(xys2)
	if err != nil {
		log.Fatal(err)
	}
	line2.LineStyle.Color = color.Gray{Y: 192}

	p.Add(plotter.NewGrid())
	p.Add(line)
	p.Add(line2)
	p.Legend.Add("f", line)
	p.Legend.Add("df/dx", line2)

	if err := p.Save(8*vg.Inch, 4*vg.Inch, "formula.png"); err != nil {
		log.Fatal(err)
	}
}
