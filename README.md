# kmeans

k-means clustering algorithm written in Go

## What It Does

[k-means clustering](https://en.wikipedia.org/wiki/K-means_clustering) partitions
an n-dimensional data set into k clusters, where each data point belongs to the
cluster with the nearest mean, serving as a prototype of the cluster.

![kmeans animation](https://github.com/muesli/kmeans/blob/master/kmeans.gif)

## Example

```go
import "github.com/muesli/kmeans"

km := kmeans.New()

// set up "random" data points (float64 values between 0.0 and 1.0)
d := []kmeans.Point{}
for x := 0; x < 255; x += 4 {
	for y := 0; y < 255; y += 4 {
		d = append(d, kmeans.Point{
			float64(x) / 255.0,
			float64(y) / 255.0,
		})
	}
}

// Partition the data points into 16 clusters
clusters, err := km.Partition(d, 16)

for _, c := range clusters {
	fmt.Printf("Centered at x: %.2f y: %.2f\n", c.Center[0]*255.0, c.Center[1]*255.0)
	fmt.Printf("Points: %+v\n\n", c.Points)
}
```

## Development

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/muesli/kmeans)
[![Build Status](https://travis-ci.org/muesli/kmeans.svg?branch=master)](https://travis-ci.org/muesli/kmeans)
[![Go ReportCard](http://goreportcard.com/badge/muesli/kmeans)](http://goreportcard.com/report/muesli/kmeans)
