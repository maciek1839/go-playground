package tasks

import (
	"fmt"
	"golang.org/x/exp/slog"
	"sync"
)

func Channels() {
	// https://medium.com/@ninucium/go-interview-questions-part-1-pointers-channels-and-range-67c61345cf3c
	slog.Info("")
	slog.Info("======> Channels")

	//
	// What's wrong with it?
	//
	//
	//ch := make(chan *int, 4)
	//array := []int{1, 2, 3, 4}
	//wg := sync.WaitGroup{}
	//wg.Add(1)
	//go func() {
	//	for _, value := range array {
	//		ch <- &value
	//	}
	//}()
	//go func() {
	//	for value := range ch {
	//		fmt.Println(*value) // what will be printed here?
	//	}
	//	wg.Done()
	//}()
	//
	//wg.Wait()

	//
	//
	//
	//
	//
	//
	// When dealing with goroutines, it makes sense to immediately analyze the code for the possibility of deadlock.
	//
	// On the one hand, everything looks correct:
	//
	// - We wait for all goroutines to finish using WaitGroup and wg.Wait()
	// - We increase the wait group counter by 1 using wg.Add(1) before calling the first goroutine
	// - In the second goroutine, we decrease it by 1 after reading all data from the ch channel wg.Done()
	//
	// But here it is necessary to remember how range works with reading data from a channel.
	// The point is that the ch channel is not closed,
	// that’s why the line wg.Done() in the second goroutine will never be reached.
	// The for value := range ch loop will be blocked waiting for the channel
	// to be closed if no values are written to the channel for reading.
	//
	// Therefore, the answer is: we will have a deadlock!
	//
	// How to fix the deadlock?
	// Deadlock can be fixed in several ways.
	// The simplest and correct way is to close the channel after all writes to the channel in the first goroutine.
	problem1Solution()
	problem1AlternativeSolution()

	//
	// Now that our code is working, what will be outputted by fmt.Println(*value)?
	//
	// It will output 4, 4, 4, 4.
	//
	// For optimization purposes in Go, new variables for the index and value will not be created for each iteration.
	// Instead, two variables are created that will be used to store the index and value in each iteration of the range loop.
	//
	// In the buffered channel ch <- &value, we pass the same pointer, and since the channel is buffered,
	// we will pass all 4 elements to it at once.
	// By the time we read the value from the channel and print the result,
	// the first range loop has already completed and the value of the last element (4) will be at the address of the value.
	// Therefore, fmt.Println(*value) will output 4, 4, 4, 4

	//
	//
	// And how can we fix this?
	//
	// To fix this, it is enough to not pass a pointer to the value directly,
	// but to create a new variable v each time and copy the value to it in each iteration,
	// and then pass its address to the channel.
	problem1LoggingSolution()

	exitGoroutine()
}

func exitGoroutine() {
	quit := make(chan bool)
	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				fmt.Println("default...")
			}
		}
	}()
	quit <- true
}

func problem1Solution() {
	slog.Info("")
	ch := make(chan *int, 4)
	array := []int{1, 2, 3, 4}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for _, value := range array {
			ch <- &value
		}
		// Close channel to reach wg.Done() line.
		close(ch)
	}()
	go func() {
		for value := range ch {
			fmt.Println("solution ", *value) // what will be printed here?
		}
		wg.Done()
	}()

	defer wg.Wait()
}

func problem1AlternativeSolution() {
	slog.Info("")
	// Here we increased the counter by the length of the array, and then in the second goroutine,
	// while iterating through the channel with wg.Done() in the loop for value := range ch,
	// we decreased the counter to 0.
	//
	// After that, the counter in WaitGroup will be reset to 0 and the program will complete execution.
	ch := make(chan *int, 4)
	array := []int{1, 2, 3, 4}
	wg := sync.WaitGroup{}
	// Add the number of works equal to the number of array elements.
	wg.Add(len(array))
	go func() {
		for _, value := range array {
			ch <- &value
		}
	}()
	go func() {
		for value := range ch {
			fmt.Println("alternative ", *value)
			// Decrement the waitgroup counter with each iteration.
			wg.Done()
		}
	}()

	wg.Wait()

	//
	// And what will happen if we don’t close the channel?
	//
	// In our case, if we don’t close the channel, nothing bad will happen,
	// since the program will immediately terminate and all memory will be freed.
	// However, it is important to remember that if a channel is not closed in a long-running application,
	// data leakage may occur and memory will be held until the program terminates.
	//
	// After making this correction, fmt.Println(*value) will output 1, 2, 3, 4.
	//
	//
	//
	//
	//
	// Let’s go back to the original code that crashed with a deadlock.
	// What will happen if we start another goroutine with an infinite for loop?
	// Here’s the code:
	problem1AnotherGoroutine()
	//
	//
	// Will there be a deadlock here?
	//
	// No. Deadlock occurs when all goroutines in the program are blocked and cannot continue their work, for example,
	// waiting for access to shared resources or waiting for each other. For example,
	// if one goroutine is blocked on reading from an empty channel,
	// and another goroutine is blocked on writing to the same channel,
	// this will lead to a deadlock.
	// Both goroutines will be waiting for each other and will not be able to continue executing their tasks.
	// If one goroutine is not blocked and all others are, this will not lead to a deadlock.
	// In this case, we added a new goroutine at the end that infinitely increments the variable a.
}

func problem1LoggingSolution() {
	slog.Info("")
	ch := make(chan *int, 4)
	array := []int{1, 2, 3, 4}
	wg := sync.WaitGroup{}
	wg.Add(len(array))
	go func() {
		for _, value := range array {
			// Temporary variable v that has new memory address every iteration.
			v := value
			ch <- &v
		}
	}()
	go func() {
		for value := range ch {
			fmt.Println(*value)
			wg.Done()
		}
	}()

	wg.Wait()
}

func problem1AnotherGoroutine() {
	ch := make(chan *int, 4)
	array := []int{1, 2, 3, 4}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for _, value := range array {
			ch <- &value
		}
	}()
	// This code below will run infinitely. It's commented out.
	// --------
	//go func() {
	//	for value := range ch {
	//		fmt.Println(*value)
	//	}
	//	wg.Done()
	//}()
	//
	//// New goroutine is run.
	//go func() {
	//	var a int
	//	for {
	//		a++
	//	}
	//}()
	//
	//wg.Wait()
	// --------
}
