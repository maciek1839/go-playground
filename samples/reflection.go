package samples

import (
	"fmt"
	"log/slog"
	"reflect"
)

func Reflection() {
	slog.Info("")
	slog.Info("======> Reflection")
	slog.Info("Reflection is the ability of a program to inspect its variables and values at run time and find their type.")
	slog.Info("Shortly speaking, reflection gives you the ability to examine types at runtime.")

	// https://blog.logrocket.com/reflection-go-use-cases-tutorial/
	x := 10
	name := "Go Lang"
	type Book struct {
		name   string
		author string
	}
	sampleBook := Book{"Reflection in Go", "John"}
	fmt.Println(reflect.TypeOf(x).Kind())          // int
	fmt.Println(reflect.TypeOf(name).Kind())       // string
	fmt.Println(reflect.TypeOf(sampleBook).Kind()) // struct
}
