package tasks

import (
	"fmt"
	"log/slog"
)

type Figure interface {
	Area() float64
}

type Rectangle struct {

	// declaring struct variables
	length float64
	width  float64
}

type Square struct {

	// declaring struct variable
	side float64
}

func (rect Rectangle) Area() float64 {

	// Area of rectangle = l * b
	area := rect.length * rect.width
	return area
}

func (sq Square) Area() float64 {

	// Area of square = a * a
	area := sq.side * sq.side
	return area
}

func Polymorphism() {
	// https://www.geeksforgeeks.org/polymorphism-in-golang/
	slog.Info("")
	slog.Info("======> Polymorphism")

	rectangle := Rectangle{

		length: 10.5,
		width:  12.25,
	}

	square := Square{

		side: 15.0,
	}

	var f1 Figure = rectangle
	var f2 Figure = square

	// printing the calculated result
	fmt.Printf("Area of rectangle: %.3f unit sq.\n", f1.Area())
	fmt.Printf("Area of square: %.3f unit sq.\n", f2.Area())
}
