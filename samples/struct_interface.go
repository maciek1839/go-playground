package samples

import (
	"fmt"
	"golang.org/x/exp/slog"
	"math"
)

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

type geometry interface {
	area() float64
	perim() float64
}

// methods associated with a given struct
func (r *rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// embedding

type base struct {
	num int
}

type container struct {
	base
	str string
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

func Structs() {
	slog.Info("")
	slog.Info("======> Slices")
	slog.Info("Go’s structs are typed collections of fields.")

	fmt.Println(person{"Bob", 20})
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("Jon"))

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)

	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())

	slog.Info("====> Interfaces")
	slog.Info("It is a description of all functions that an object must have in order to be an \"X\"")

	r = rect{width: 3, height: 4}
	c := circle{radius: 5}
	measure(&r)
	measure(c)

	// embedding
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	slog.Info(co.describe())
}
