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
	d := []kmeans.Point{}
	// for l := 0; l < 255; l += 16 {
	for a := 0; a < 255; a += 4 {
		for b := 0; b < 255; b += 4 {
			d = append(d, kmeans.Point{
				// float64(l) / 255.0,
				float64(a) / 255.0,
				float64(b) / 255.0,
			})
		}
	}
	// }
	fmt.Printf("%d data points\n", len(d))

	km, err := kmeans.NewWithOptions(0.01, true)
	if err != nil {
		panic(err)
	}
	clusters, err := km.Run(d, 16)
	if err != nil {
		panic(err)
	}

	buffer := bytes.NewBuffer([]byte{})
	buffer.Write([]byte(header))

	for i, c := range clusters {
		fmt.Printf("Cluster: %d %+v\n", i, c.Center)
		col := colorful.Lab(80, -0.9+(c.Center[0]*1.8), -0.9+(c.Center[1]*1.8))
		// LAB: col := colorful.Lab(c.Center[0]*1.0, -0.9+(c.Center[1]*1.8), -0.9+(c.Center[2]*1.8))
		// RGB: col := colorful.Color{R: c.Center[0], G: c.Center[1], B: c.Center[2]}
		fmt.Println("Hex:", col.Hex())

		buffer.Write([]byte(fmt.Sprintf(cell, col.Hex())))
	}

	buffer.Write([]byte(footer))
	if err := ioutil.WriteFile("palette.html", buffer.Bytes(), 0644); err != nil {
		panic(err)
	}
}
