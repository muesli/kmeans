package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/muesli/clusters"
	"github.com/muesli/kmeans"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// set up a random two-dimensional data set (float64 values between 0.0 and 1.0)
	var d clusters.Observations
	for x := 0; x < 1024; x++ {
		d = append(d, clusters.Coordinates{
			rand.Float64(),
			rand.Float64(),
		})
	}
	fmt.Printf("%d data points\n", len(d))

	// Partition the data points into 7 clusters
	km := kmeans.New()
	clusters, _ := km.Partition(d, 7)

	for i, c := range clusters {
		fmt.Printf("Cluster: %d\n", i)
		fmt.Printf("Centered at x: %.2f y: %.2f\n", c.Center[0], c.Center[1])
	}
}
