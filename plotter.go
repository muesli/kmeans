package kmeans

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/muesli/clusters"

	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

// The Plotter interface lets you implement your own plotters
type Plotter interface {
	Plot(cc clusters.Clusters, iteration int) error
}

// SimplePlotter is the default standard plotter for 2-dimensional data sets
type SimplePlotter struct {
}

// A monokai-ish color palette
var colors = []drawing.Color{
	drawing.ColorFromHex("f92672"),
	drawing.ColorFromHex("89bdff"),
	drawing.ColorFromHex("66d9ef"),
	drawing.ColorFromHex("67210c"),
	drawing.ColorFromHex("7acd10"),
	drawing.ColorFromHex("af619f"),
	drawing.ColorFromHex("fd971f"),
	drawing.ColorFromHex("dcc060"),
	drawing.ColorFromHex("545250"),
	drawing.ColorFromHex("4b7509"),
}

// Plot draw a 2-dimensional data set into a PNG file named {k_iteration}.png
func (p SimplePlotter) Plot(cc clusters.Clusters, iteration int) error {
	var series []chart.Series

	// draw data points
	for i, c := range cc {
		series = append(series, chart.ContinuousSeries{
			Style: chart.Style{
				Show:        true,
				StrokeWidth: chart.Disabled,
				DotColor:    colors[i%len(colors)],
				DotWidth:    8,
			},
			XValues: c.PointsInDimension(0),
			YValues: c.PointsInDimension(1),
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
		XValues: cc.CentersInDimension(0),
		YValues: cc.CentersInDimension(1),
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
		return err
	}
	return ioutil.WriteFile(fmt.Sprintf("%d_%d.png", len(cc), iteration), buffer.Bytes(), 0644)
}
