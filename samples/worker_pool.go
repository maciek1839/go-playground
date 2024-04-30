package samples

import (
	"fmt"
	"golang.org/x/exp/slog"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func WorkerPools() {
	// https://gobyexample.com/worker-pools
	// Another approach: https://doziestar.medium.com/effortlessly-tame-concurrency-in-golang-a-deep-dive-into-worker-pools-31f6ed9f2872
	slog.Info("")
	slog.Info("======> Worker pools")
	slog.Info("A worker pool is a concurrency pattern made up of a set number of worker goroutines that are in charge of executing tasks concurrently.")

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
