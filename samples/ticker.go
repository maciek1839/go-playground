package samples

import (
	"fmt"
	"golang.org/x/exp/slog"
	"time"
)

func Tickers() {
	slog.Info("")
	slog.Info("======> Tickers")
	slog.Info("\nTimers are for when you want to do something once in the future - tickers are for when you want to do something repeatedly at regular intervals.")

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
