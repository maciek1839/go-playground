package tsp

import (
	"math"
	"sync"
)

// SolveTSPConcurrentBruteForce evaluates all permutations concurrently (exact solution)
func SolveTSPConcurrentBruteForce(distance [][]float64) ([]int, float64) {
	n := len(distance)
	if n == 0 {
		return nil, 0
	}

	cities := make([]int, n)
	for i := 0; i < n; i++ {
		cities[i] = i
	}

	perms := permute(cities[1:]) // fix city 0
	var wg sync.WaitGroup
	resultChan := make(chan struct {
		route    []int
		distance float64
	}, len(perms))

	for _, perm := range perms {
		wg.Add(1)
		go func(p []int) {
			defer wg.Done()
			route := append([]int{0}, p...)
			route = append(route, 0)
			total := totalDistance(route, distance)
			resultChan <- struct {
				route    []int
				distance float64
			}{route, total}
		}(perm)
	}

	wg.Wait()
	close(resultChan)

	shortest := math.MaxFloat64
	var bestRoute []int
	for res := range resultChan {
		if res.distance < shortest {
			shortest = res.distance
			bestRoute = res.route
		}
	}

	return bestRoute, shortest
}

// SolveTSPConcurrentGreedy runs greedy TSP from each city concurrently and returns the best
func SolveTSPConcurrentGreedy(distance [][]float64) ([]int, float64) {
	n := len(distance)
	if n == 0 {
		return nil, 0
	}

	var wg sync.WaitGroup
	resultChan := make(chan struct {
		route    []int
		distance float64
	}, n)

	for start := 0; start < n; start++ {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			visited := make([]bool, n)
			route := []int{start}
			visited[start] = true
			total := 0.0
			current := start

			for len(route) < n {
				next := -1
				shortest := math.MaxFloat64
				for i := 0; i < n; i++ {
					if !visited[i] && distance[current][i] < shortest {
						shortest = distance[current][i]
						next = i
					}
				}
				if next == -1 {
					break
				}
				visited[next] = true
				route = append(route, next)
				total += shortest
				current = next
			}

			total += distance[current][start]
			route = append(route, start)

			resultChan <- struct {
				route    []int
				distance float64
			}{route, total}
		}(start)
	}

	wg.Wait()
	close(resultChan)

	shortest := math.MaxFloat64
	var bestRoute []int
	for res := range resultChan {
		if res.distance < shortest {
			shortest = res.distance
			bestRoute = res.route
		}
	}

	return bestRoute, shortest
}
