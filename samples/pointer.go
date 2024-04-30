package samples

import (
	"golang.org/x/exp/slog"
	"strconv"
)

type Person struct {
	Name string
	Age  int
}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func Pointers() {
	slog.Info("")
	slog.Info("======> Pointers")
	slog.Info("Pointers are a powerful concept in Go programming that allows you to work directly with memory addresses,")
	slog.Info("enabling you to manipulate data more efficiently.")
	slog.Info("A pointer is a variable that holds the memory address of another variable.")

	var num int = 42
	var ptr *int
	ptr = &num
	slog.Info("Value of 'num': " + strconv.Itoa(*ptr))

	var p *Person = &Person{"John", 30}
	slog.Info("Name: " + (*p).Name)
	slog.Info("Age: " + strconv.Itoa(p.Age))

	i := 1
	slog.Info("initial:" + strconv.Itoa(i))
	zeroval(i)
	slog.Info("zeroval:" + strconv.Itoa(i))

	zeroptr(&i)
	slog.Info("zeroptr:" + strconv.Itoa(i))

}
