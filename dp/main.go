package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {

	scatterData := fourPoints(4)
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "MinerMap"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	p.Add(plotter.NewGrid())

	s, err := plotter.NewScatter(scatterData)
	if err != nil {
		panic(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	p.Add(s)

	edgeData := edge()
	sedge, err := plotter.NewScatter(edgeData)
	if err != nil {
		panic(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 0}
	p.Add(sedge)

	if err := p.Save(7*vg.Inch, 7*vg.Inch, "scatter.png"); err != nil {
		panic(err)
	}

}

func fourPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	pts[0].X = -7.0
	pts[0].Y = -7.0
	pts[1].X = 7.0
	pts[1].Y = -7.0

	pts[2].X = 7.0
	pts[2].Y = 7.0
	pts[3].X = -7.0
	pts[3].Y = 7.0

	return pts
}

func edge() plotter.XYs {

	pts := make(plotter.XYs, 16*4)
	ptsindex := 0
	for i := -7; i <= 7; i++ {
		pts[ptsindex].X = (float64)(i)
		pts[ptsindex].Y = (float64)(7)
		ptsindex++
		pts[ptsindex].X = (float64)(i)
		pts[ptsindex].Y = (float64)(-7)
		ptsindex++
		pts[ptsindex].X = (float64)(-7)
		pts[ptsindex].Y = (float64)(i)
		ptsindex++
		pts[ptsindex].X = (float64)(7)
		pts[ptsindex].Y = (float64)(i)
		ptsindex++
	}
	return pts
}
