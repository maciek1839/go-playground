package samples

import (
	"fmt"
	"golang.org/x/exp/slog"
)

func mayPanic() {
	panic("a problem")
}
func Recover() {
	slog.Info("")
	slog.Info("======> Recover")
	slog.Info("Go makes it possible to recover from a panic, by using the recover built-in function.")
	slog.Info("recover must be called within a deferred function.")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()
	mayPanic()
	fmt.Println("After mayPanic()")
}
