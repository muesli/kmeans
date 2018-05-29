package kmeans

import (
	"fmt"
	"math"
)

// Point is a data point (float64 between 0.0 and 1.0) in n dimensions
type Point []float64

// Points is a slice of points
type Points []Point

// Distance returns the euclidean distance between two data points
func (p Point) Distance(p2 Point) float64 {
	var r float64
	for i, v := range p {
		r += math.Pow(v-p2[i], 2)
	}
	return r
}

// Equal returns true if the two points have equal values
func (p Point) Equal(p2 Point) bool {
	for i := range p {
		if p[i] != p2[i] {
			return false
		}
	}

	return true
}

// Mean returns the mean point of p
func (p Points) Mean() (Point, error) {
	var l = len(p)
	if l == 0 {
		return Point{}, fmt.Errorf("there is no mean for an empty set of points")
	}

	c := make([]float64, len(p[0]))
	for _, point := range p {
		for j, v := range point {
			c[j] += v
		}
	}

	var point Point
	for _, v := range c {
		point = append(point, v/float64(l))
	}
	return point, nil
}
