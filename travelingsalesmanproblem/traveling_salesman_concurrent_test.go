package tsp

import (
	"testing"
)

var testMatrix = [][]float64{
	{0, 10, 15, 20},
	{10, 0, 35, 25},
	{15, 35, 0, 30},
	{20, 25, 30, 0},
}

func TestSolveTSPConcurrentBruteForce(t *testing.T) {
	route, dist := SolveTSPConcurrentBruteForce(testMatrix)
	expected := 80.0

	if dist != expected {
		t.Errorf("Concurrent Brute Force: expected %.2f, got %.2f", expected, dist)
	}
	if len(route) != 5 {
		t.Errorf("Concurrent Brute Force: expected route length 5, got %d", len(route))
	}
}

func TestSolveTSPConcurrentGreedy(t *testing.T) {
	route, dist := SolveTSPConcurrentGreedy(testMatrix)
	if len(route) != 5 {
		t.Errorf("Concurrent Greedy: expected route length 5, got %d", len(route))
	}
	if dist < 70 {
		t.Errorf("Concurrent Greedy: suspiciously low distance %.2f", dist)
	}
}
