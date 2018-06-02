# kmeans

k-means clustering algorithm implementation written in Go

## What It Does

[k-means clustering](https://en.wikipedia.org/wiki/K-means_clustering) partitions
a multi-dimensional data set into `k` clusters, where each data point belongs
to the cluster with the nearest mean, serving as a prototype of the cluster.

![kmeans animation](https://github.com/muesli/kmeans/blob/master/kmeans.gif)

## When Should I Use It?

- When you have numeric, multi-dimensional data sets
- You don't have labels for your data
- You know exactly how many clusters you want to partition your data into

## Example

```go
import (
	"github.com/muesli/kmeans"
	"github.com/muesli/clusters"
)

// set up a random two-dimensional data set (float64 values between 0.0 and 1.0)
var d clusters.Observations
for x := 0; x < 1024; x++ {
	d = append(d, clusters.Coordinates{
		rand.Float64(),
		rand.Float64(),
	})
}

// Partition the data points into 16 clusters
km := kmeans.New()
clusters, err := km.Partition(d, 16)

for _, c := range clusters {
	fmt.Printf("Centered at x: %.2f y: %.2f\n", c.Center[0], c.Center[1])
	fmt.Printf("Matching data points: %+v\n\n", c.Observations)
}
```

## Complexity

If `k` (the amount of clusters) and `d` (the dimensions) are fixed, the problem
can be exactly solved in time O(n<sup>dk+1</sup>), where `n` is the number of
entities to be clustered.

The running time of the algorithm is O(nkdi), where `n` is the number of
`d`-dimensional vectors, `k` the number of clusters and `i` the number of
iterations needed until convergence. On data that does have a clustering
structure, the number of iterations until convergence is often small, and
results only improve slightly after the first dozen iterations. The algorithm
is therefore often considered to be of "linear" complexity in practice,
although it is in the worst case superpolynomial when performed until
convergence.

## Options

You can greatly reduce the running time by adjusting the required delta
threshold. With the following options the algorithm finishes when less than 5%
of the data points shifted their cluster assignment in the last iteration:

```go
km, err := kmeans.NewWithOptions(0.05, nil)
```

The default setting for the delta threshold is 0.01 (1%).

If you are working with two-dimensional data sets, kmeans can generate
beautiful graphs (like the one above) for each iteration of the algorithm:

```go
km, err := kmeans.NewWithOptions(0.01, kmeans.SimplePlotter{})
```

Careful: this will generate PNGs in your current working directory.

You can write your own plotters by implementing the `kmeans.Plotter` interface.

## Development

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/muesli/kmeans)
[![Build Status](https://travis-ci.org/muesli/kmeans.svg?branch=master)](https://travis-ci.org/muesli/kmeans)
[![Coverage Status](https://coveralls.io/repos/github/muesli/kmeans/badge.svg?branch=master)](https://coveralls.io/github/muesli/kmeans?branch=master)
[![Go ReportCard](http://goreportcard.com/badge/muesli/kmeans)](http://goreportcard.com/report/muesli/kmeans)
