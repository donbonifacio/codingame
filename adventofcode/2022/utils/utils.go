package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {}

func Atoi(raw string) int {
	val, err := strconv.Atoi(raw)
	if err != nil {
		panic(err)
	}
	return val
}

func AsLines(data string) []string {
	return Split(data, "\n")
}

func Split(data string, separator string) []string {
	parts := strings.Split(strings.TrimSpace(data), separator)
	for i := 0; i < len(parts); i++ {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

func IsEmpty(data string) bool {
	return strings.TrimSpace(data) == ""
}

func EachString(data []string, op func(str string)) {
	for i := 0; i < len(data); i++ {
		item := data[i]
		op(item)
	}
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func ReadInput(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func ReadFileWithNumbers(fileName string) []int {
	raw, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	rawString := string(raw)
	parts := strings.Split(rawString, ",")
	numbers := []int{}

	for i := 0; i < len(parts); i++ {
		part := strings.TrimSpace(parts[i])
		num, err := strconv.Atoi(part)

		if err != nil {
			panic(fmt.Sprintf("Error parsing number '%v'", num))
		}

		numbers = append(numbers, num)
	}

	return numbers
}

func WriteFile(fileName string, content string) {
	f, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(content)

	if err2 != nil {
		panic(err)
	}
}
