package tasks

import (
	"bytes"
	"fmt"
	"golang.org/x/exp/slog"
)

func Slices() {
	// https://medium.com/@ninucium/go-interview-questions-part-2-slices-87f5289fb7eb
	slog.Info("")
	slog.Info("======> Slices")

	referenceType()
	sliceExpansion()
	sliceExpansion2()
	subSlices()

	x := []int{3, 2, 1}
	reverse(x)

	checkIfSliceIsEmpty()
	compareSlices()
}

func compareSlices() {
	x1 := []byte{'C', 'O', 'D', 'I', 'N', 'G'}
	x2 := []byte{'N', 'I', 'N', 'J', 'A', 'S'}
	output := bytes.Compare(x1, x2)
	if output == 0 {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not Equal")
	}
}

func checkIfSliceIsEmpty() {
	// https://www.educative.io/blog/50-golang-interview-questions
	// Create a program that checks if a slice is empty. Find the simplest solution.
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	// The easiest way to check if a slice is empty is to use the built-in len() function, which returns the length of a slice.
	// If len(slice) == 0, then you know the slice is empty.
	r := [3]int{1, 2, 3}

	if len(r) == 0 {
		fmt.Println("Empty! ", r)

	} else {
		fmt.Println("Not Empty!", r)
	}
}

func reverse(sw []int) {
	// https://www.educative.io/blog/50-golang-interview-questions
	// Reverse the order of a slice
	// Implement function reverse that takes a slice of integers and reverses the slice in place without using a temporary slice.
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	for a, b := 0, len(sw)-1; a < b; a, b = a+1, b-1 {
		sw[a], sw[b] = sw[b], sw[a]
	}
}

func sliceExpansion2() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	//
	//
	//
	//
	// SUMMARY
	// [John Paul George Ringo]
	// [John Paul] [Paul George]
	// [John XXX] [XXX George]
	// [John XXX George Ringo]
}

func referenceType() {
	slog.Info("====> Reference Type")
	var x []int
	// x=[1], len=1, cap=1
	x = append(x, 1)
	// When we use the append function to add a new element, it internally checks the slice for capacity.
	// Since in our case the new slice length will become larger than the capacity of underlying array,
	// a slice extension will be performed.
	// The new one with the capacity: cap = current cap * 2
	// This rule for increasing capacity applies to a slice length less than 1024.
	// Beyond that, the slice will not be increased by 100% (doubled), but by 25%.
	// x=[1, 2], len=2, cap=2
	x = append(x, 2)
	// x=[1,2,3], len=3, cap=4
	x = append(x, 3)
	// We create a new variable y, which equals our slice x: y=[1,2,3], len=3, cap=4
	y := x
	// x=[1,2,3,4], len=4, cap=4
	x = append(x, 4)
	// We created y when x was equal to x=[1,2,3], len=3, cap=4.
	// In y we keep a reference to the first element of the slice, our array length is 3 and capacity is 4.
	// y=[1,2,3,5], len=4, cap=4
	y = append(y, 5)
	// In the previous steps with append, we did not have slice expansion.
	// Therefore, both x and y point to the same element at index 0.
	// The operation x[0] = 0 will put a new value 0 in the element in both array x and array y.
	x[0] = 0

	fmt.Println(x)
	fmt.Println(y)

	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	// SUMMARY
	//  var x []int      // x=[], len=0, cap=0
	//  x = append(x, 1) // x=[1], len=1, cap=1
	//  x = append(x, 2) // x=[1, 2], len=2, cap=2
	//  x = append(x, 3) // x=[1, 2, 3], len=3, cap=4
	//  y := x           // y=[1, 2, 3], len=3, cap=4
	//  x = append(x, 4) // x=[1, 2, 3, 4], len=4, cap=4
	//  y = append(y, 5) // x=[1, 2, 3, 5], len=4, cap=4
	//  x[0] = 0         // x=[0, 2, 3, 5], len=4, cap=4
	//
	//  fmt.Println(x)   // x=[0, 2, 3, 5], len=4, cap=4
	//  fmt.Println(y)   // y=[0, 2, 3, 5], len=4, cap=4

	// LONG SUMMARY
	// var x []int
	// x=[1], len=1, cap=1
	// x = append(x, 1)
	// When we use the append function to add a new element, it internally checks the slice for capacity.
	// Since in our case the new slice length will become larger than the capacity of underlying array,
	// a slice extension will be performed.
	// The new one with the capacity: cap = current cap * 2
	// This rule for increasing capacity applies to a slice length less than 1024.
	// Beyond that, the slice will not be increased by 100% (doubled), but by 25%.
	// x=[1, 2], len=2, cap=2
	// x = append(x, 2)
	// x=[1,2,3], len=3, cap=4
	// x = append(x, 3)
	// We create a new variable y, which equals our slice x: y=[1,2,3], len=3, cap=4
	// y := x
	// x=[1,2,3,4], len=4, cap=4
	// x = append(x, 4)
	// We created y when x was equal to x=[1,2,3], len=3, cap=4.
	// In y we keep a reference to the first element of the slice, our array length is 3 and capacity is 4.
	// y=[1,2,3,5], len=4, cap=4
	// y = append(y, 5)
	// In the previous steps with append, we did not have slice expansion.
	// Therefore, both x and y point to the same element at index 0.
	// The operation x[0] = 0 will put a new value 0 in the element in both array x and array y.
	// x[0] = 0
}

func sliceExpansion() {
	slog.Info("====> Slice Expansion")
	x := []int{1, 2, 3, 4}
	y := x
	x = append(x, 5)
	y = append(y, 6)
	x[0] = 0
	fmt.Println(x)
	fmt.Println(y)

	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	// SUMMARY
	//  x := []int{1,2,3,4} // x=[1,2,3,4], len=4, cap=4
	//  y := x              // y=[1,2,3,4], len=4, cap=4
	//  x = append(x, 5)    // x=[1,2,3,4,5], len=5, cap=8
	//  y = append(y, 6)    // y=[1,2,3,4,6], len=5, cap=8
	//  x[0] = 0            // x=[0,2,3,4,5], len=5, cap=8
	//
	//  fmt.Println(x)      // x=[0,2,3,4,5], len=5, cap=8
	//  fmt.Println(y)      // y=[1,2,3,4,6], len=5, cap=8

	// LONG SUMMARY
	//	x := []int{1, 2, 3, 4}
	//	y := x
	// During the x = append(x, 5) the slice expanded into a new area of memory,
	// where the capacity now became equal to 8.
	//
	// Now x and y stopped referring to the same first element of the slice, and ended up in different areas of memory.
	// Similarly, for this reason adding the value 6 to the slice y did not affect the x slice at all.
	//	x = append(x, 5)
	//	y = append(y, 6)
	//	x[0] = 0
	//	fmt.Println(x)
	//	fmt.Println(y)
}

func subSlices() {
	slog.Info("====> Sub Slices")
	x := []int{1, 2, 3, 4, 5}
	x = append(x, 6)
	x = append(x, 7)
	a := x[4:]
	y := alterSlice(a)
	fmt.Println(a)
	fmt.Println(x)
	fmt.Println(y)

	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	// SUMMARY
	// x := []int{1, 2, 3, 4, 5} // [1,2,3,4,5], len=5, cap=5
	// x = append(x, 6)          // [1,2,3,4,5,6], len=6, cap=10
	// x = append(x, 7)          // [1,2,3,4,5,6,7], len=7, cap=10
	// a := x[4:]                // [5,6,7], len=3, cap=6
	// y := alterSlice(a)        // [10, 6, 7, 11], len=4, cap=6
	//
	// fmt.Println(a)            // [10 6 7]
	// fmt.Println(x)            // [1,2,3,4,10,6,7], len=7, cap=10
	// fmt.Println(y)            // [10, 6, 7, 11]
	//
	// func alterSlice(a []int) []int {
	//	a[0] = 10                 // [10, 6, 7], len=3, cap=6
	//	a = append(a, 11)         // [10, 6, 7, 11], len=4, cap=6
	//	return a
	// }
}

func alterSlice(a []int) []int {
	a[0] = 10
	a = append(a, 11)
	return a
}
