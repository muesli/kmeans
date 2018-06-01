// Package kmeans implements the k-means clustering algorithm
// See: https://en.wikipedia.org/wiki/K-means_clustering
package kmeans

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Kmeans configuration/option struct
type Kmeans struct {
	// when a plotter is set, Plot gets called after each iteration
	plotter Plotter
	// deltaThreshold (in percent between 0.0 and 0.1) aborts processing if
	// less than n% of data points shifted clusters in the last iteration
	deltaThreshold float64
	// iterationThreshold aborts processing when the specified amount of
	// algorithm iterations was reached
	iterationThreshold int
}

// NewWithOptions returns a Kmeans configuration struct with custom settings
func NewWithOptions(deltaThreshold float64, plotter Plotter) (Kmeans, error) {
	if deltaThreshold <= 0.0 || deltaThreshold >= 1.0 {
		return Kmeans{}, fmt.Errorf("threshold is out of bounds (must be >0.0 and <1.0, in percent)")
	}

	return Kmeans{
		plotter:            plotter,
		deltaThreshold:     deltaThreshold,
		iterationThreshold: 96,
	}, nil
}

// New returns a Kmeans configuration struct with default settings
func New() Kmeans {
	m, _ := NewWithOptions(0.01, nil)
	return m
}

func randomizeClusters(k int, dataset Points) (Clusters, error) {
	var c Clusters
	if len(dataset) == 0 || len(dataset[0]) == 0 {
		return c, fmt.Errorf("there must be at least one dimension in the data set")
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
	if k > len(dataset) {
		return Clusters{}, fmt.Errorf("the size of the data set must at least equal k")
	}

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
			ci := clusters.Nearest(point)
			clusters[ci].Points = append(clusters[ci].Points, point)
			if points[p] != ci {
				points[p] = ci
				changes++
			}
		}

		for ci := 0; ci < len(clusters); ci++ {
			if len(clusters[ci].Points) == 0 {
				// During the iterations, if any of the cluster centers has no
				// data points associated with it, assign a random data point
				// to it.
				// Also see: http://user.ceng.metu.edu.tr/~tcan/ceng465_f1314/Schedule/KMeansEmpty.html
				var ri int
				for {
					// find a cluster with at least two data points, otherwise
					// we're just emptying one cluster to fill another
					ri = rand.Intn(len(dataset))
					if len(clusters[points[ri]].Points) > 1 {
						break
					}
				}
				clusters[ci].Points = append(clusters[ci].Points, dataset[ri])
				points[ri] = ci
			}
		}

		if changes > 0 {
			clusters.recenter()
		}
		if m.plotter != nil {
			m.plotter.Plot(clusters, i)
		}
		if i == m.iterationThreshold ||
			changes < int(float64(len(dataset))*m.deltaThreshold) {
			// fmt.Println("Aborting:", changes, int(float64(len(dataset))*m.TerminationThreshold))
			break
		}
	}

	return clusters, nil
}

func (m Kmeans) Sil(dataset Points) (int, float64, error) {
	var d float64
	var mc int

	for n := 2; n < 10; n++ {
		cc, err := m.Partition(dataset, n)
		if err != nil {
			return 0, -1.0, err
		}

		var si float64
		var sc int64
		for ci, c := range cc {
			for _, p := range c.Points {
				ai := p.averageDistance(c.Points) // FIXME: exclude p
				_, bi := cc.Dissimilarity(p, ci)

				si += bi - ai/math.Max(ai, bi)
				sc++
			}
		}

		sd := si / float64(sc)
		fmt.Println("sil:", sd)
		if mc == 0 || sd < d {
			mc = n
			d = sd
		}
	}

	return mc, d, nil
}
