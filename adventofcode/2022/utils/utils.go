package utils

import (
	"strconv"
)

func main() {}

func Atoi(raw string) int {
	val, err := strconv.Atoi(raw)
	if err != nil {
		panic(err)
	}
	return val
}
