package samples

import (
	"fmt"
	"golang.org/x/exp/slog"
	"strconv"
)

type ExampleStruct struct {
	string1 string
	var1    float32
}

func Variables() {
	slog.Info("")
	slog.Info("======> Variables")
	slog.Info("Variables are the names you give to computer memory locations which are used to store values in a computer program")

	var a = "string"
	var b string
	slog.Info("String variables: a= " + a + " b= " + b)

	var c int8
	var d float32 = 231.322
	slog.Info(
		"Number variables: c= " +
			strconv.FormatInt(int64(c), 10) +
			" d= " + strconv.FormatFloat(float64(d), 'f', -1, 64),
	)

	// If a variable is not assigned any value, Go automatically initializes it with the zero value of the variable's type.
	var e ExampleStruct
	slog.Info("Struct: " + fmt.Sprintf("%#v", &e))
	slog.Info("Struct.string1 " + e.string1 + " Struct.var1 " + strconv.Itoa(int(e.var1)))

	const f = 100_000
	slog.Info("f: " + strconv.Itoa(f))
}
