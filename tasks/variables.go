package tasks

import (
	"fmt"
	"log/slog"
)

func Variables() {
	// https://www.educative.io/blog/50-golang-interview-questions
	slog.Info("")
	slog.Info("======> Variables")
	//
	//
	// Swap the values of two variables without a temporary variable.
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	// SOLUTION
	//
	fmt.Println(swap())
}

func swap() []int {
	a, b := 15, 10
	b, a = a, b
	return []int{a, b}
}
