package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/muesli/clusters"
	"github.com/muesli/kmeans"

	colorful "github.com/lucasb-eyer/go-colorful"
)

var (
	header = `
	<html>
	<body>
		<table width="100%" height="100%">
			<tr>
	`

	cell = `
				<td bgcolor="%s" />
	`

	footer = `
			</tr>
		</table>
	</body>
	</html>
	`
)

type Color struct {
	colorful.Color
}

func (c Color) Coordinates() clusters.Coordinates {
	l, a, b := c.Lab()
	return clusters.Coordinates{l, a, b}
}

func (c Color) Distance(pos clusters.Coordinates) float64 {
	c2 := colorful.Lab(pos[0], pos[1], pos[2])
	return c.DistanceLab(c2)
}

func main() {
	// Create data points in the CIE L*a*b color space
	// l for lightness channel
	// a, b for color channels
	var d clusters.Observations
	for l := 0.2; l < 0.8; l += 0.05 {
		for a := -1.0; a < 1.0; a += 0.1 {
			for b := -1.0; b < 1.0; b += 0.1 {
				c := colorful.Lab(l, a, b)
				if !c.IsValid() {
					continue
				}

				d = append(d, Color{c})
			}
		}
	}

	// Write HTML header
	buffer := bytes.NewBuffer([]byte{})
	buffer.Write([]byte(header))

	// Enable graph generation (.png files) for each iteration
	km, _ := kmeans.NewWithOptions(0.01, kmeans.SimplePlotter{})

	// Partition the color space into 16 clusters (palette colors)
	clusters, _ := km.Partition(d, 16)

	for i, c := range clusters {
		fmt.Printf("Cluster: %d %+v\n", i, c.Center)
		col := colorful.Lab(c.Center[0], c.Center[1], c.Center[2]).Clamped()
		fmt.Println("Color as Hex:", col.Hex())

		buffer.Write([]byte(fmt.Sprintf(cell, col.Hex())))
	}

	// Write HTML footer and generate palette.html
	buffer.Write([]byte(footer))
	ioutil.WriteFile("palette.html", buffer.Bytes(), 0644)
}
