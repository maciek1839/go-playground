package samples

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slog"
)

func MapKeys[K constraints.Ordered, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}
func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

// another example

func Max[T constraints.Ordered](slice []T) T {
	if len(slice) == 0 {
		panic("empty slice")
	}
	max := slice[0]
	for _, value := range slice[1:] {
		if value > max {
			max = value
		}
	}
	return max
}

func Generics() {
	slog.Info("")
	slog.Info("======> Generics")
	slog.Info("Generic programming is a style of computer programming in which algorithms are written in terms of types to-be-specified-later that are then instantiated when needed for specific types provided as parameters.")

	slice := []int{989, 1, 4, 177, 5, 90, 2, 6, 3, 1001, 475, 233}
	max := Max(slice)
	fmt.Println(max)

	var m = map[int]string{1: "2", 2: "4", 4: "8"}
	fmt.Println("keys:", MapKeys(m))

	_ = MapKeys[int, string](m)
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}
