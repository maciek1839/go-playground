package samples

import (
	"golang.org/x/exp/slog"
	"strconv"
)

func Functions() {
	slog.Info("")
	slog.Info("======> Functions")
	slog.Info("Functions are blocks of reusable and organised code that usually perform a single, related action.")

	variadicFunctions()

	closure()
}

func closure() {
	slog.Info("====> Closures")
	slog.Info("Go supports anonymous functions, which can form closures.")
	slog.Info("Anonymous functions are useful when you want to define a function inline without having to name it.")

	// The anonymous function closes over the variable i to form a closure.
	const N = 100
	j := func() int {
		return N + 1
	}()
	slog.Info("Closure: " + strconv.Itoa(j))
}

func variadicFunctions() {
	slog.Info("====> Variadic functions")
	slog.Info("Variadic functions can be called with any number of trailing arguments.")

	nums := []int{1, 2, 3, 4}
	slog.Info("Sum: " + strconv.Itoa(sum(nums...)))
}

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}
