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
	Debug bool
	// DeltaThreshold (in percent between 0.0 and 0.1) aborts processing if
	// less than n% of data points shifted clusters in the last iteration
	DeltaThreshold float64
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

// Run executes the k-means algorithm on the given dataset and divides it into k clusters
func (m Kmeans) Run(dataset Points, k int) (Clusters, error) {
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
		if m.Debug {
			draw(clusters, strconv.Itoa(i))
		}

		if changes < int(float64(len(dataset))*m.DeltaThreshold) {
			// fmt.Println("Aborting:", changes, int(float64(len(dataset))*m.TerminationThreshold))
			break
		}
	}

	return clusters, nil
}
