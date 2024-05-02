package tasks

import (
	"fmt"
	"log/slog"
)

func SumOfSquares() {
	// https://www.educative.io/blog/50-golang-interview-questions
	slog.Info("")
	slog.Info("======> Sum of squares")
	// Implement the SumOfSquares function which takes an integer, c and returns the sum of all squares between 1 and c.
	// You’ll need to use select statements, goroutines, and channels.
	mychannel := make(chan int)
	quitchannel := make(chan int)
	sum := 0
	go func() {
		for i := 1; i <= 5; i++ {
			sum += <-mychannel
		}
		fmt.Println(sum)
		quitchannel <- 0
	}()
	sumOfSquares(mychannel, quitchannel)
}

func sumOfSquares(c, quit chan int) {
	y := 1
	for {
		select {
		case c <- (y * y):
			y++
		case <-quit:
			return
		}
	}
}
