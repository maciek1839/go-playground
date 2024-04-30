package samples

import (
	"fmt"
	"golang.org/x/exp/slog"
)

func Loops() {
	slog.Info("")
	slog.Info("======> Loops")

	i := 1
	for i <= 3 {
		fmt.Print("for1 ", i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Print("standard for ", j)
	}

	var arr = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i := range arr {
		fmt.Print("range ", i)
	}

	for {
		fmt.Println("loop")
		break
	}
}
