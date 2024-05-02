package tasks

import "fmt"

func Shadowing() {
	x := 10
	fmt.Println("Outer x:", x) // Outer x: 10

	{
		x := 20
		fmt.Println("Inner x:", x) // Inner x: 20
	}

	fmt.Println("Outer x:", x) // Outer x: 10
}
