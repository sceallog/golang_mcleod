package main

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	numbers := strings.Split(in, " ")
	max, _ := strconv.Atoi(numbers[0])
	min, _ := strconv.Atoi(numbers[0])

	for _, numStr := range numbers {
		num, _ := strconv.Atoi(numStr)
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	return fmt.Sprintf("%d %d", max, min)
}

func main() {
	fmt.Println(HighAndLow("1 9 3 4 -5"))
}

// func ReplaceAll(s, old, new string) string
