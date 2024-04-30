package samples

import (
	"fmt"
	"golang.org/x/exp/slog"
	"regexp"
)

func Regex() {
	slog.Info("")
	slog.Info("======> Regex")
	slog.Info("A regular expression sometimes referred to as rational expression, is a sequence of characters that specifies a match pattern in text.")

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString("peach punch"))
}
