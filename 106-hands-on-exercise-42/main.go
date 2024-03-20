package main

import "fmt"

func main() {
	numbers := [5]int{}
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5

	for _, v := range numbers {
		fmt.Printf("Value: %v\t Type: %T\n", v, v)
	}
}
