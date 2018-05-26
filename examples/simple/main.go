package main

import (
	"fmt"

	"github.com/muesli/kmeans"
)

func main() {
	// set up "random" data points (float64 values between 0.0 and 1.0)
	km := kmeans.New()
	d := []kmeans.Point{}
	for x := 0; x < 255; x += 4 {
		for y := 0; y < 255; y += 4 {
			d = append(d, kmeans.Point{
				float64(x) / 255.0,
				float64(y) / 255.0,
			})
		}
	}
	fmt.Printf("%d data points\n", len(d))

	// Partition the data points into 4 clusters
	clusters, _ := km.Partition(d, 4)

	for i, c := range clusters {
		fmt.Printf("Cluster: %d\n", i)
		fmt.Printf("Centered at x: %.2f y: %.2f\n", c.Center[0]*255.0, c.Center[1]*255.0)
	}
}
