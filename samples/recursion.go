package samples

import (
	"golang.org/x/exp/slog"
	"strconv"
)

func Recursions() {
	slog.Info("")
	slog.Info("======> Recursion")
	slog.Info("Recursion is defined as a process which calls itself directly or indirectly and the corresponding function is called a recursive function.")
	slog.Info("Recursion is a method of solving a computational problem where the solution depends on solutions to smaller instances of the same problem.")

	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	slog.Info("Recursive function (fib): " + strconv.Itoa(fib(7)))
	slog.Info("Recursive function (fact): " + strconv.Itoa(fact(7)))
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
