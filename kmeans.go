// Package kmeans implements the k-means clustering algorithm
// See: https://en.wikipedia.org/wiki/K-means_clustering
package kmeans

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Kmeans configuration/option struct
type Kmeans struct {
	// when Debug is enabled, graphs are generated after each iteration
	debug bool
	// DeltaThreshold (in percent between 0.0 and 0.1) aborts processing if
	// less than n% of data points shifted clusters in the last iteration
	deltaThreshold float64
}

// NewWithOptions returns a Kmeans configuration struct with custom settings
func NewWithOptions(deltaThreshold float64, debug bool) (Kmeans, error) {
	if deltaThreshold <= 0.0 || deltaThreshold >= 1.0 {
		return Kmeans{}, fmt.Errorf("threshold is out of bounds (must be >0.0 and <1.0, in percent)")
	}

	return Kmeans{
		debug:          debug,
		deltaThreshold: deltaThreshold,
	}, nil
}

// New returns a Kmeans configuration struct with default settings
func New() Kmeans {
	m, _ := NewWithOptions(0.01, false)
	return m
}

func randomizeClusters(k int, dataset Points) (Clusters, error) {
	var c Clusters
	if len(dataset) == 0 || len(dataset[0]) == 0 {
		return c, fmt.Errorf("there must be at least one dimension in the data-set")
	}
	if k == 0 {
		return c, fmt.Errorf("k must be greater than 0")
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < k; i++ {
		var p Point
		for j := 0; j < len(dataset[0]); j++ {
			p = append(p, rand.Float64())
		}

		c = append(c, Cluster{
			Center: p,
		})
	}
	return c, nil
}

// Partition executes the k-means algorithm on the given dataset and
// partitions it into k clusters
func (m Kmeans) Partition(dataset Points, k int) (Clusters, error) {
	clusters, err := randomizeClusters(k, dataset)
	if err != nil {
		return Clusters{}, err
	}

	points := make([]int, len(dataset))
	changes := 1

	for i := 0; changes > 0; i++ {
		changes = 0
		clusters.reset()

		for p, point := range dataset {
			ci := clusters.nearestCluster(point)
			clusters[ci].Points = append(clusters[ci].Points, point)
			if points[p] != ci {
				points[p] = ci
				changes++
			}
		}

		for ci, c := range clusters {
			if len(c.Points) == 0 {
				// During the iterations, if any of the cluster centers has no
				// data points associated with it, assign a random data point
				// to it.
				//
				// Also see: http://user.ceng.metu.edu.tr/~tcan/ceng465_f1314/Schedule/KMeansEmpty.html
				ri := rand.Intn(len(dataset))
				c.Points = append(c.Points, dataset[ri])

				// FIXME: remove Point from previously assigned cluster?
				points[ri] = ci
			}
		}

		if changes > 0 {
			clusters.recenter()
		}
		if m.debug {
			draw(clusters, strconv.Itoa(i))
		}

		if changes < int(float64(len(dataset))*m.deltaThreshold) {
			// fmt.Println("Aborting:", changes, int(float64(len(dataset))*m.TerminationThreshold))
			break
		}
	}

	return clusters, nil
}
