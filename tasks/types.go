package tasks

import (
	"fmt"
	"log/slog"
)

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Double %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know  type %T!\n", v)
	}
}

func Types() {
	// https://www.educative.io/blog/50-golang-interview-questions
	slog.Info("")
	slog.Info("======> Types")

	do(21)
	do("hello")
	do(true)
}
