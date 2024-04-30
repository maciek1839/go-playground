package samples

import (
	"fmt"
	"golang.org/x/exp/slog"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

func Concurrency() {
	slog.Info("======> Concurrency")
	slog.Info("Concurrency is the ability of different parts or units of a program to be executed out-of-order or in partial order, without affecting the outcome.")
	slog.Info("")

	// https://betterprogramming.pub/golang-how-to-implement-concurrency-with-goroutines-channels-2b78b8077984
	slog.Info("===> Goroutines and Channels")
	slog.Info("")

	goroutine()
	channel()
	nonBlockingChannel()
	channelDirections()
	channelSynchronization()
	timeouts()
	waitGroup()
	rateLimiting()
	atomicCounter()
}

func atomicCounter() {
	var ops atomic.Uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {

				ops.Add(1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops:", ops.Load())
}

func rateLimiting() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.NewTicker(200 * time.Millisecond)

	for req := range requests {
		<-limiter.C
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

func channelSynchronization() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// We await the worker using the synchronization approach
	<-done

	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}

func nonBlockingChannel() {
	// Basic sends and receives on channels are blocking. However,
	// we can use select with a default clause to implement non-blocking sends, receives,
	// and even non-blocking multi-way selects.
	// https://gobyexample.com/non-blocking-channel-operations

	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

func timeouts() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func channelDirections() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func worker1(r chan string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		r <- fmt.Sprintf("1.%d", i)
	}

	wg.Done()
}

func worker2(r chan string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		r <- fmt.Sprintf("2.%d", i)
	}

	wg.Done()
}

func waitGroup() {
	slog.Info("")
	slog.Info("==> WaitGroup")
	slog.Info("WaitGroup is used to wait for all the goroutines launched here to finish.")

	var wg sync.WaitGroup

	r := make(chan string)
	wg.Add(2)
	go worker1(r, &wg)
	go worker2(r, &wg)

	go func() {
		defer close(r)
		wg.Wait()
	}()

	for i := range r {
		slog.Info("Got job result: " + i)
	}
}

func channel() {
	slog.Info("")
	slog.Info("==> Channels")

	// Create an unbuffered channel of type int
	c := make(chan int)

	// Start a Goroutine that sends a value on the channel
	go func() {
		slog.Info("A simple goroutine starts...")
		time.Sleep(2 * time.Second)
		c <- 82
	}()

	// Receive the value from the channel
	// When the main program executes the goroutines, it waits for the channel to get some data before continuing.
	value := <-c

	slog.Info("Received value: " + strconv.Itoa(value))

	slog.Info("")
	slog.Info("=> Buffered channels")
	slog.Info("If we need the goroutine to return multiple values, we have to use buffered channels.")

	multiply := func(arr []int, ch chan int) {
		for _, elem := range arr {
			slog.Info("Multiplying...")
			ch <- elem * 3
		}
	}

	arr := []int{2, 3, 4}
	ch := make(chan int, len(arr))
	go multiply(arr, ch)

	for i := 0; i < len(arr); i++ {
		// use fmt.Printf to distinguish from normal logs
		slog.Info("Result: " + strconv.Itoa(<-ch))
	}

	arr2 := []int{2, 3, 4}
	ch2 := make(chan int, len(arr))
	// anonymous function
	go func(arr []int, ch chan int) {
		for _, elem := range arr {
			ch <- elem * 3
		}
	}(arr2, ch2)
	for i := 0; i < len(arr); i++ {
		slog.Info("Result2: " + strconv.Itoa(<-ch2))
	}

	slog.Info("")
	slog.Info("=> Channels between goroutines")
	slog.Info("Channels not only work for interactions between a goroutine and the main programs, they also provide a way to communicate between different goroutine.")

	arr3 := []int{2, 3, 4}
	ch3 := make(chan int, len(arr))
	go func(arr []int, ch chan int) {
		minusCh := make(chan int, 3)
		for _, elem := range arr {
			valueTmp := elem * 3
			if valueTmp%2 == 0 {
				go minusThree(valueTmp, minusCh)
				valueTmp = <-minusCh
			}
			ch <- valueTmp
		}
	}(arr3, ch3)
	for i := 0; i < len(arr3); i++ {
		slog.Info("Result3: " + strconv.Itoa(<-ch3))
	}

	slog.Info("")
	slog.Info("=> Range and close")
	slog.Info("With the instruction, for i := range ch we can iterate over the goroutine’s results as soon as they are sent. The goroutine should close the channel with the function close once it finishes sending data.")

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		slog.Info(elem)
	}

	slog.Info(Multiline(
		"You needn’t close every channel when you’ve finished with it.",
		"It’s only necessary to close a channel when it is important to tell the receiving goroutines that all data have been sent."))

	slog.Info("")
	slog.Info("=> Select")
	slog.Info(Multiline(
		"We can read from multiple channels at the same time",
		"Works as a way to wait for multiple channels at the same time, preventing one from blocking another.",
	))

	arr4 := []int{2, 3, 4, 5, 6}
	ch4 := make(chan int, len(arr4))
	minusCh := make(chan int, len(arr4))

	go func(arr []int, ch chan int) {
		for _, elem := range arr {
			ch <- elem * 3
		}
	}(arr4, ch4)
	go func(arr []int, ch chan int) {
		for _, elem := range arr {
			ch <- elem - 3
		}
	}(arr4, ch4)

	for i := 0; i < len(arr4)*2; i++ {
		select {
		case msg1 := <-ch4:
			fmt.Printf("Result timesThree: %v \n", msg1)
		case msg2 := <-minusCh:
			fmt.Printf("Result minusThree: %v \n", msg2)
		default:
			slog.Info("Non blocking way of listening to multiple channels")
		}
	}

	slog.Info("")
	slog.Info("=> Mutual exclusion")
	var n = 1
	var mu sync.Mutex
	for i := 0; i < 2; i++ {
		go func(num int) {
			mu.Lock()
			defer mu.Unlock()
			num *= 3
			slog.Info("Mutex value: " + strconv.Itoa(num))
		}(n)
	}
	time.Sleep(time.Second)
}

func minusThree(number int, ch chan int) {
	ch <- number - 3
	slog.Info("The functions continues after returning the result")
}

func goroutine() {
	// https://futurice.com/blog/gocurrency
	slog.Info("==> Goroutines")

	car1 := "Ferrari"
	car2 := "Lamborghini"

	// Create a Goroutine for each car
	go race(car1)
	go race(car2)

	// Wait for the race to finish
	time.Sleep(5 * time.Second)

	slog.Info("The race is over!")
}

func race(car string) {
	for i := 0; i < 5; i++ {
		slog.Info(car + " is racing...")
		time.Sleep(time.Second)
	}
}
