package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// draw output
func draw() {
	wg.Add(1)
	defer wg.Done()
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "MinerMap"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	p.Add(plotter.NewGrid())
	edgeData := edge()
	sedge, err := plotter.NewScatter(edgeData)
	if err != nil {
		panic(err)
	}
	sedge.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 0}
	p.Add(sedge)
	pts := make(plotter.XYs, 1024)
	ptsindex := 0
	for {
		pd, ok := <-drawpoint
		if ptsindex == 1024 {
			addPoints(pts, p)
			ptsindex = 0
		}
		pts[ptsindex].X = (float64)(pd.X)
		pts[ptsindex].Y = (float64)(pd.Y)
		ptsindex++
		if ok == false {
			pts = pts[0:ptsindex]
			addPoints(pts, p)
			if err := p.Save((vg.Length)(xylimit)*vg.Inch, (vg.Length)(xylimit)*vg.Inch, "miner.png"); err != nil {
				panic(err)
			}
			return
		}
	}
}

func addPoints(pts plotter.XYs, p *plot.Plot) {
	plotData, err := plotter.NewScatter(pts)
	if err != nil {
		panic(err)
	}
	plotData.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	p.Add(plotData)
}

func edge() plotter.XYs {
	len := 4 * (2*xylimit + 1)
	pts := make(plotter.XYs, len)
	ptsindex := 0
	for i := xylimit * (-1); i <= xylimit; i++ {
		pts[ptsindex].X = (float64)(i)
		pts[ptsindex].Y = (float64)(xylimit)
		ptsindex++
		pts[ptsindex].X = (float64)(i)
		pts[ptsindex].Y = (float64)(xylimit * (-1))
		ptsindex++
		pts[ptsindex].X = (float64)(xylimit * (-1))
		pts[ptsindex].Y = (float64)(i)
		ptsindex++
		pts[ptsindex].X = (float64)(xylimit)
		pts[ptsindex].Y = (float64)(i)
		ptsindex++
	}
	return pts
}
