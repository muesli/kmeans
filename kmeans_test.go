package kmeans

import (
	"testing"
)

func TestNewErrors(t *testing.T) {
	_, err := NewWithOptions(0.00, nil)
	if err == nil {
		t.Errorf("Expected invalid options to return an error, got nil")
	}

	_, err = NewWithOptions(1.00, nil)
	if err == nil {
		t.Errorf("Expected invalid options to return an error, got nil")
	}
}

func TestPartitioningError(t *testing.T) {
	km := New()
	if _, err := km.Partition(Points{}, 1); err == nil {
		t.Errorf("Expected error partitioning with empty data set, got nil")
		return
	}

	d := Points{
		Point{
			0.1,
			0.1,
		},
	}
	if _, err := km.Partition(d, 0); err == nil {
		t.Errorf("Expected error partitioning with 0 clusters, got nil")
		return
	}

	if _, err := km.Partition(d, 2); err == nil {
		t.Errorf("Expected error partitioning with more clusters than data points, got nil")
		return
	}
}

func TestDimensions(t *testing.T) {
	var d Points
	for x := 0; x < 255; x += 32 {
		for y := 0; y < 255; y += 32 {
			d = append(d, Point{
				float64(x) / 255.0,
				float64(y) / 255.0,
			})
		}
	}

	k := 4
	km := New()
	clusters, err := km.Partition(d, k)
	if err != nil {
		t.Errorf("Unexpected error partitioning: %v", err)
		return
	}

	if len(clusters) != k {
		t.Errorf("Expected %d clusters, got: %d", k, len(clusters))
	}
}
