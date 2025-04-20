package tsp

/**
The Traveling Salesman Problem is a classic optimization problem in computer science and operations research.
The goal is to find the shortest possible route that visits each city exactly once and returns to the starting city.

Given a list of cities and the distances between them, TSP asks:
- What is the shortest possible route that visits each city exactly once and returns to the origin city?

This problem is NP-hard, meaning that for large numbers of cities,
there's no known efficient way to find an exact solution quickly.
However, for small sets of cities, we can solve it exactly using brute-force or dynamic programming.


In this solution, we:
- Represent cities and distances using a distance matrix.
- Use a brute-force approach (by generating all permutations of cities).
- Calculate the total distance for each permutation.
- Track the shortest one.

This is suitable for a small number of cities (e.g. ≤10) due to the factorial time complexity.

Brute Force
How it works:
- Generates every possible route (permutation) between cities.
- Calculates the total distance for each one.
- Picks the shortest route from all possibilities.

Pros:
- Always gives the optimal (shortest possible) path.

Cons:
- Extremely slow: Time complexity is O(n!) — grows very fast with more cities.
- Not practical for more than ~10 cities.

0 → 1 → 2 → 3 → 0
0 → 1 → 3 → 2 → 0
0 → 2 → 1 → 3 → 0
...and all others


Greedy (Nearest Neighbor)
How it works:

- Start at city 0.
- At each step, go to the closest unvisited city.
- Continue until all cities are visited.
- Return to the starting city.

Pros:
- Very fast: O(n²) time.
- Easy to implement.

Cons:
- Doesn’t always give the shortest route.
- Can get stuck in a suboptimal path early.

Example: Start at city 0:
Nearest is city 1 → go there.
Then nearest is city 3 → go there.
Then city 2 → go there.
Return to 0.

This route might not be optimal.
*/

import (
	"math"
)

// SolveTSPBruteForce finds the shortest route using brute-force (all permutations)
func SolveTSPBruteForce(distance [][]float64) ([]int, float64) {
	n := len(distance)
	if n == 0 {
		return nil, 0
	}

	cities := make([]int, n)
	for i := 0; i < n; i++ {
		cities[i] = i
	}

	minDistance := math.MaxFloat64
	var bestRoute []int

	permutations := permute(cities[1:]) // fix city 0 as start

	for _, perm := range permutations {
		route := append([]int{0}, perm...)
		route = append(route, 0) // return to start
		total := totalDistance(route, distance)
		if total < minDistance {
			minDistance = total
			bestRoute = make([]int, len(route))
			copy(bestRoute, route)
		}
	}
	return bestRoute, minDistance
}

// SolveTSPGreedy finds a quick route using the Nearest Neighbor heuristic
func SolveTSPGreedy(distance [][]float64) ([]int, float64) {
	n := len(distance)
	if n == 0 {
		return nil, 0
	}

	visited := make([]bool, n)
	route := []int{0}
	visited[0] = true
	total := 0.0
	current := 0

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
		route = append(route, next)
		visited[next] = true
		total += shortest
		current = next
	}

	// Return to start
	total += distance[current][0]
	route = append(route, 0)

	return route, total
}

// Helpers below are kept unexported (lowercase)
func totalDistance(route []int, distance [][]float64) float64 {
	sum := 0.0
	for i := 0; i < len(route)-1; i++ {
		sum += distance[route[i]][route[i+1]]
	}
	return sum
}

func permute(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(a []int, i int) {
		if i == len(a)-1 {
			perm := make([]int, len(a))
			copy(perm, a)
			res = append(res, perm)
			return
		}
		for j := i; j < len(a); j++ {
			a[i], a[j] = a[j], a[i]
			helper(a, i+1)
			a[i], a[j] = a[j], a[i]
		}
	}
	helper(arr, 0)
	return res
}
