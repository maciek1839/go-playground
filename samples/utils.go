package samples

import (
	"fmt"
	"golang.org/x/exp/slog"
	"reflect"
	"strconv"
	"strings"
)

func Multiline(parts ...string) string {
	return strings.Join(parts, "\n")
}

func ArrayToString[T any](arr []T) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(arr)), ","), "[]")
}

func PrintArrayInfo[T any](arr []T) {
	slog.Info("length: " + strconv.Itoa(len(arr)) + " capacity: " + strconv.Itoa(cap(arr)) + " kind: " + reflect.ValueOf(arr).Kind().String())
}
