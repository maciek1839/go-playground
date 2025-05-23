package samples

import (
	"fmt"
	"golang.org/x/exp/slog"
)

func Arrays() {
	slog.Info("")
	slog.Info("======> Arrays")
	slog.Info("An array is a numbered sequence of elements of a specific length.")
	slog.Info("In typical Go code, slices are much more common; arrays are useful in some special scenarios.")

	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)

	testArray := []string{"one", "two", "three", "four", "five"}
	array1(testArray)
	fmt.Println("testArray: ", testArray)
}

func array1(arr []string) {
	arr[0] = "change"
}
