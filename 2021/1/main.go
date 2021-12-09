package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input")
	input := strings.Split(string(content), "\n")
	counter := 0
	for index, item := range input {
		if index != 0 && index < len(input) {
			last, _ := strconv.Atoi(input[index-1])
			actual, _ := strconv.Atoi(item)
			counter += isBiggerThanTheLastNumber(last, actual)
		}
	}
	fmt.Printf("%d numbers bigger ", counter)
}

func isBiggerThanTheLastNumber(last, actual int) int {
	if actual > last {
		return 1
	}
	return 0
}
