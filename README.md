# kmeans

[k-means clustering algorithm](https://en.wikipedia.org/wiki/K-means_clustering) written in Go

## What It Does

k-means clustering partitions an n-dimensional data set into k clusters, where
each data point belongs to the cluster with the nearest mean, serving as a
prototype of the cluster.

![kmeans animation](https://github.com/muesli/kmeans/blob/master/kmeans.gif)

## Example

```
import "github.com/muesli/kmeans"

km := kmeans.New()

// set up a 2d grid with "random" data points between 0.0 and 1.0
d := []kmeans.Point{}
for x := 0; x < 255; x += 4 {
	for y := 0; y < 255; y += 4 {
		d = append(d, kmeans.Point{
			float64(y) / 255.0,
			float64(y) / 255.0,
		})
	}
}

// Partition the data points into 4 clusters
clusters, err := km.Run(d, 4)

for _, c := range clusters {
	fmt.Printf("Center: X: %.2f Y: %.2f\n", c.Center[0]*255.0, c.Center[1]*255.0)
	fmt.Printf("Points: %+v\n\n", c.Points)
}
```

## Development

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/muesli/kmeans)
[![Build Status](https://travis-ci.org/muesli/kmeans.svg?branch=master)](https://travis-ci.org/muesli/kmeans)
[![Go ReportCard](http://goreportcard.com/badge/muesli/kmeans)](http://goreportcard.com/report/muesli/kmeans)
