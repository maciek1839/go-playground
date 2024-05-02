package tasks

import (
	"fmt"
	"log/slog"
)

func Stack() {
	// https://www.educative.io/blog/50-golang-interview-questions
	slog.Info("")
	slog.Info("======> Stack")

	//
	// TASK: Implement a Stack (LIFO)
	//
	//
	//
	//
	// SOLUTION: You can implement a stack using a slice object.
	var stack []string

	// Push
	stack = append(stack, "world!")
	stack = append(stack, "Hello ")

	for len(stack) > 0 {

		// Print top
		n := len(stack) - 1

		fmt.Print(stack[n])

		// Pop
		stack = stack[:n]
	}
}
