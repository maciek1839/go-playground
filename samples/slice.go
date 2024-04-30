package samples

import (
	"cmp"
	"fmt"
	"golang.org/x/exp/slices"
	"golang.org/x/exp/slog"
	"strconv"
)

func Slices() {
	slog.Info("")
	slog.Info("======> Slices")
	slog.Info("Slices are similar to arrays, but are more powerful and flexible.")
	slog.Info("Slices are pointers to arrays.")
	slog.Info("An array is a fixed-size data structure that stores a series of values of the same type, while a slice is a variable-size data structure that references a portion of an array.")

	slice0 := make([]string, 3)
	slog.Info(ArrayToString(slice0))

	slice1 := []int{1, 2, 3}
	slog.Info("length: " + strconv.Itoa(len(slice1)) + "capacity: " + strconv.Itoa(cap(slice1)))
	slog.Info(ArrayToString(slice1))

	slice1copy := make([]int, len(slice1))
	copy(slice1copy, slice1)
	slog.Info("slice1 copy: " + ArrayToString(slice1copy))

	slice2 := []string{"Go", "Slices", "Are", "Powerful"}
	slog.Info("length: " + strconv.Itoa(len(slice2)) + " capacity: " + strconv.Itoa(cap(slice2)))
	slog.Info(ArrayToString(slice2))

	// new returns a pointer to the memory allocated,
	// while make returns the value of the type being created.
	// new only works with basic types such as int, float, bool, etc.
	// make is used for creating slices, maps, and channels.
	//
	// new allocates zeroed memory, while make allocates memory and initializes it.
	slice3 := new([]int32)
	*slice3 = append(*slice3, 10, 20, 30)
	PrintArrayInfo(*slice3)

	slice4 := make([]int32, 0)
	slog.Info("value: " + fmt.Sprint(slice4) + " address: " + fmt.Sprintf("%p", slice4))
	slice4 = append(slice4, 10, 20, 30)
	// Append will only allocate a new array if there isn't sufficient capacity in the slice you're appending to.
	// If you need to have a separate array, use make to create a new slice and use copy to copy whatever you need from the original slice.
	// https://stackoverflow.com/questions/43991754/why-golang-append-same-slice-result-will-share-one-memory-address
	slog.Info("value: " + fmt.Sprint(slice4) + " address: " + fmt.Sprintf("%p", slice4))
	PrintArrayInfo(slice4)

	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s = append(s, "d")
	s = append(s, "e", "f")

	l := s[2:5]
	slog.Info("sl1:" + ArrayToString(l))
	l = s[:5]
	slog.Info("sl2:" + ArrayToString(l))
	fmt.Println("sl2:", l)
	l = s[2:]
	slog.Info("sl3:" + ArrayToString(l))
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	slog.Info("dcl:" + ArrayToString(t))
	t2 := []string{"g", "h", "i"}
	// Two array values are equal if their corresponding elements are equal.
	if slices.Equal(t, t2) {
		slog.Info("t == t2")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	sorting()
	sortingFunction()
}

func sorting() {
	slog.Info("Sorting")
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:   ", ints)

	sorted := slices.IsSorted(ints)
	fmt.Println("Sorted: ", sorted)
}

func sortingFunction() {
	fruits := []string{"peach", "banana", "kiwi"}

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	type Person struct {
		name string
		age  int
	}
	people := []Person{
		Person{name: "Jax", age: 37},
		Person{name: "TJ", age: 25},
		Person{name: "Alex", age: 72},
	}

	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		})
	fmt.Println(people)
}
