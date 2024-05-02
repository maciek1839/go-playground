package tasks

import (
	"fmt"
	"log/slog"
)

func MinMax() {
	// https://www.educative.io/blog/50-golang-interview-questions
	slog.Info("")
	slog.Info("======> Min/Max")
	//
	//
	// Implement Min(x, y int) and Max(x, y int) functions that take two integers and return the lesser or greater value, respectively.
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
	// By default, Go only supports min and max for floats using math.min and math.max.
	// You’ll have to create your own implementations to make it work for integers.
	// Min returns the smallest of x or y.
	fmt.Println(Min(5, 10))
	fmt.Println(Max(5, 10))
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// Max returns the larger of x or y.

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
