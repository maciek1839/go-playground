package samples

import (
	"fmt"
	"golang.org/x/exp/slog"
	"time"
)

func Timers() {
	slog.Info("")
	slog.Info("======> Timers")
	slog.Info("Timers represent a single event in the future.")

	timer1 := time.NewTimer(1 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(time.Second)
}
