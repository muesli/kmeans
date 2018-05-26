package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

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

func main() {
	// Create data points in the CIE L*a*b color space
	d := []kmeans.Point{}

	// l for lightness channel
	// a, b for color channels
	for l := 30; l < 255; l += 16 {
		for a := 0; a < 255; a += 16 {
			for b := 0; b < 255; b += 16 {
				d = append(d, kmeans.Point{
					float64(l) / 255.0,
					float64(a) / 255.0,
					float64(b) / 255.0,
				})
			}
		}
	}

	// Write HTML header
	buffer := bytes.NewBuffer([]byte{})
	buffer.Write([]byte(header))

	// Enable graph generation (.png files) for each iteration
	km, _ := kmeans.NewWithOptions(0.01, true)

	// Partition the color space into 16 clusters (palette colors)
	clusters, _ := km.Partition(d, 16)

	for i, c := range clusters {
		fmt.Printf("Cluster: %d %+v\n", i, c.Center)
		col := colorful.Lab(c.Center[0]*1.0, -0.9+(c.Center[1]*1.8), -0.9+(c.Center[2]*1.8))
		fmt.Println("Color as Hex:", col.Hex())

		buffer.Write([]byte(fmt.Sprintf(cell, col.Hex())))
	}

	// Write HTML footer and generate palette.html
	buffer.Write([]byte(footer))
	ioutil.WriteFile("palette.html", buffer.Bytes(), 0644)
}
