package tasks

func Map() {
	// https://www.naukri.com/code360/library/golang-interview-questions
	//
	// You can copy the values of a Map variable in GoLang by traversing its keys.
	// The simplest way to copy a map in GoLang is given below.
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
	a := map[string]int{"A": 1, "B": 2}
	b := make(map[string]int)
	for key, value := range a {
		b[key] = value
	}
}
