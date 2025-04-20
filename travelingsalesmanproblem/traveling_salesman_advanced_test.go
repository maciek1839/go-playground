package tsp

import (
	"testing"
)

// Test SolveTSPGenetic
func TestSolveTSPGenetic(t *testing.T) {
	dist := [][]float64{
		{0, 2, 9, 10},
		{1, 0, 6, 4},
		{15, 7, 0, 8},
		{6, 3, 12, 0},
	}
	path, cost := SolveTSPGenetic(dist, 100, 10)

	// Check if the cost is greater than 0
	if cost <= 0 {
		t.Errorf("Expected cost to be positive, got: %.2f", cost)
	}

	// Check if the path length is correct
	if len(path) != len(dist) {
		t.Errorf("Expected path length %d, got: %d", len(dist), len(path))
	}

	// Check if the path does not revisit any city
	visited := make(map[int]bool)
	for _, city := range path {
		if visited[city] {
			t.Errorf("City %d was visited more than once", city)
		}
		visited[city] = true
	}
}

// Test SolveTSPAnnealing
func TestSolveTSPAnnealing(t *testing.T) {
	dist := [][]float64{
		{0, 2, 9, 10},
		{1, 0, 6, 4},
		{15, 7, 0, 8},
		{6, 3, 12, 0},
	}
	path, cost := SolveTSPAnnealing(dist, 1000, 0.995, 1000)

	// Check if the cost is greater than 0
	if cost <= 0 {
		t.Errorf("Expected cost to be positive, got: %.2f", cost)
	}

	// Check if the path length is correct
	if len(path) != len(dist) {
		t.Errorf("Expected path length %d, got: %d", len(dist), len(path))
	}

	// Check if the path does not revisit any city
	visited := make(map[int]bool)
	for _, city := range path {
		if visited[city] {
			t.Errorf("City %d was visited more than once", city)
		}
		visited[city] = true
	}
}
