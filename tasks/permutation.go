package tasks

import (
	"fmt"
	"log/slog"
)

func Permutation() {
	// https://www.educative.io/blog/50-golang-interview-questions
	slog.Info("")
	slog.Info("======> Permutation")
	//
	//
	// Implement the perm() function that accepts a slice or string and prints all possible combinations of characters.
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
	// SOLUTION
	//
	Perm([]rune("abc"), func(a []rune) {
		fmt.Println(string(a))
	})
}

// Perm calls f with each permutation of a.
func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}
