package tasks

import "fmt"

func FibonacciSequence() {
	// https://www.naukri.com/code360/library/golang-interview-questions
	fmt.Println(Fib(10))
}

func Fib(x int) int {
	if x < 2 {
		return x
	}
	return Fib(x-1) + Fib(x-2)
}
