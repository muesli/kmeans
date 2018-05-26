package kmeans

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

var colors = []drawing.Color{
	drawing.ColorFromHex("f92672"),
	drawing.ColorFromHex("89bdff"),
	drawing.ColorFromHex("66d9ef"),
	drawing.ColorFromHex("4b7509"),
	drawing.ColorFromHex("545250"),
	drawing.ColorFromHex("67210c"),
	drawing.ColorFromHex("7acd10"),
	drawing.ColorFromHex("af619f"),
	drawing.ColorFromHex("fd971f"),
	drawing.ColorFromHex("dcc060"),
}

func (cluster *Cluster) pointInDimension(n int) Point {
	var v Point
	for _, p := range cluster.Points {
		v = append(v, p[n])
	}
	return v
}

func (clusters Clusters) centerInDimension(n int) Point {
	var v Point
	for _, c := range clusters {
		v = append(v, c.Center[n])
	}
	return v
}

func draw(clusters Clusters, iteration int) {
	var series []chart.Series

	// draw data points
	for i, c := range clusters {
		/*
			col := colorful.Lab(c.Center[0]*1.0, -0.9+(c.Center[1]*1.8), -0.9+(c.Center[2]*1.8))
			// RGB: col := colorful.Color{R: c.Center[0], G: c.Center[1], B: c.Center[2]}
			colHex := drawing.ColorFromHex(col.Hex()[1:])
		*/

		series = append(series, chart.ContinuousSeries{
			Style: chart.Style{
				Show:        true,
				StrokeWidth: chart.Disabled,
				DotColor:    colors[i%len(colors)],
				DotWidth:    8},
			XValues: c.pointInDimension(0),
			YValues: c.pointInDimension(1),
		})
	}

	// draw cluster center points
	series = append(series, chart.ContinuousSeries{
		Style: chart.Style{
			Show:        true,
			StrokeWidth: chart.Disabled,
			DotColor:    drawing.ColorBlack,
			DotWidth:    16,
		},
		XValues: clusters.centerInDimension(0),
		YValues: clusters.centerInDimension(1),
	})

	graph := chart.Chart{
		Height: 1024,
		Width:  1024,
		Series: series,
		XAxis: chart.XAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true,
			},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fmt.Sprintf("%d.png", iteration), buffer.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}
