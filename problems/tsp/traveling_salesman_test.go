package problems

import (
	"testing"
)

var sampleMatrix = [][]float64{
	{0, 10, 15, 20},
	{10, 0, 35, 25},
	{15, 35, 0, 30},
	{20, 25, 30, 0},
}

func TestSolveTSPBruteForce(t *testing.T) {
	route, dist := SolveTSPBruteForce(sampleMatrix)
	expectedDistance := 80.0

	if dist != expectedDistance {
		t.Errorf("Brute force: expected distance %.2f, got %.2f", expectedDistance, dist)
	}

	if len(route) != 5 {
		t.Errorf("Brute force: expected route of length 5, got %d", len(route))
	}
}

func TestSolveTSPGreedy(t *testing.T) {
	route, dist := SolveTSPGreedy(sampleMatrix)

	if len(route) != 5 {
		t.Errorf("Greedy: expected route of length 5, got %d", len(route))
	}

	if dist < 80.0 {
		t.Errorf("Greedy: distance %.2f seems too low (shorter than optimal?)", dist)
	}
}
