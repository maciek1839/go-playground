package samples

import (
	"fmt"
	"golang.org/x/exp/slog"
)

func Defer() {
	slog.Info("")
	slog.Info("======> Defer")
	slog.Info("Defer is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup.")

	defer fmt.Println("world")

	fmt.Println("hello")
}
