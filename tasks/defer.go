package tasks

import "fmt"

func Defer() {
	// https://go.dev/tour/flowcontrol/13
	fmt.Println("Calculating...")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done")

	//
	// SUMMARY
	// LIFO, ang. Last In, First Out
	// Deferred function calls are pushed onto a stack.
	// When a function returns, its deferred calls are executed in last-in-first-out order.
}
