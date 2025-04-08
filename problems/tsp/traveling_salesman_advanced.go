package problems

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

/**
For production or larger datasets, you'd want to use heuristics like:
- Genetic Algorithms
- Simulated Annealing
- Ant Colony Optimization
- Dynamic Programming (Held-Karp Algorithm)
*/

type Individual struct {
	path []int
	// In the context of the genetic algorithm for solving the Traveling Salesman Problem (TSP), fitness is a measure of how good a solution (i.e., a path) is.
	// Specifically, it represents the total distance of the path, with shorter paths having higher fitness.
	// The goal of the algorithm is to evolve a population of paths towards finding the shortest possible path that visits all cities exactly once and returns to the starting point.
	fitness float64
}

// SolveTSPGenetic uses a genetic algorithm to solve the Traveling Salesman Problem.
// dist is the distance matrix (2D array), generations defines the number of generations to run,
// and populationSize defines the size of the population.
// It returns the shortest path and its distance.
func SolveTSPGenetic(dist [][]float64, generations, populationSize int) ([]int, float64) {
	n := len(dist)
	randSource := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSource)

	// Create initial population
	population := make([]Individual, populationSize)
	for i := range population {
		path := randGen.Perm(n)
		population[i] = Individual{path, routeLength(path, dist)}
	}

	// GA loop
	for g := 0; g < generations; g++ {
		// Sort by fitness (ascending order)
		sortPopulation(population)

		// Keep top 50% of population
		nextGen := make([]Individual, populationSize)
		copy(nextGen[:populationSize/2], population[:populationSize/2])

		// Crossover + mutation
		for i := populationSize / 2; i < populationSize; i++ {
			p1 := population[randGen.Intn(populationSize/2)].path
			p2 := population[randGen.Intn(populationSize/2)].path
			child := crossover(p1, p2)
			mutate(child)
			nextGen[i] = Individual{child, routeLength(child, dist)}
		}

		population = nextGen
	}

	// Sort the final population
	sortPopulation(population)
	return population[0].path, population[0].fitness
}

// SolveTSPAnnealing uses the Simulated Annealing algorithm to solve the Traveling Salesman Problem
// dist is the distance matrix, initialTemp is the starting temperature for annealing,
// coolingRate is the rate at which the temperature cools, and maxIter is the maximum number of iterations.
// It returns the shortest path and its distance.
func SolveTSPAnnealing(dist [][]float64, initialTemp float64, coolingRate float64, maxIter int) ([]int, float64) {
	n := len(dist)
	randSource := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSource)
	current := randGen.Perm(n)
	currentDist := routeLength(current, dist)
	best := make([]int, n)
	copy(best, current)
	bestDist := currentDist
	temp := initialTemp

	for i := 0; i < maxIter; i++ {
		newPath := swapTwo(current)
		newDist := routeLength(newPath, dist)
		if newDist < currentDist || randGen.Float64() < math.Exp((currentDist-newDist)/temp) {
			current = newPath
			currentDist = newDist
			if newDist < bestDist {
				copy(best, newPath)
				bestDist = newDist
			}
		}
		temp *= coolingRate
	}
	return best, bestDist
}

// === Helper Functions ===
func routeLength(path []int, dist [][]float64) float64 {
	sum := 0.0
	for i := 0; i < len(path)-1; i++ {
		sum += dist[path[i]][path[i+1]]
	}
	return sum
}

// sortPopulation sorts the population by fitness (ascending order)
func sortPopulation(pop []Individual) {
	// simple insertion sort
	for i := 1; i < len(pop); i++ {
		for j := i; j > 0 && pop[j].fitness < pop[j-1].fitness; j-- {
			pop[j], pop[j-1] = pop[j-1], pop[j]
		}
	}
}

// crossover combines two parent paths into a child path
func crossover(p1, p2 []int) []int {
	n := len(p1)
	child := make([]int, 0, n)
	used := make(map[int]bool)
	for _, v := range p1[:n/2] {
		child = append(child, v)
		used[v] = true
	}
	for _, v := range p2 {
		if !used[v] {
			child = append(child, v)
		}
	}
	return child
}

// mutate randomly swaps two cities in a path
func mutate(path []int) {
	i, j := rand.Intn(len(path)), rand.Intn(len(path))
	path[i], path[j] = path[j], path[i]
}

// swapTwo swaps two cities in a path at random positions
func swapTwo(path []int) []int {
	n := len(path)
	res := make([]int, n)
	copy(res, path)
	i := rand.Intn(n)
	j := rand.Intn(n)
	res[i], res[j] = res[j], res[i]
	return res
}

// You can call PrintPath(path, dist) after any algorithm to visualize results
func PrintPath(path []int, dist [][]float64) {
	fmt.Println("Tour:")
	for i := 0; i < len(path)-1; i++ {
		fmt.Printf("%d -> ", path[i])
	}
	fmt.Printf("%d\n", path[len(path)-1])
	fmt.Printf("Total cost: %.2f\n", routeLength(path, dist))
}
