package main

import "fmt"

func main() {
	numbers := []int{}
	for i := 42; i < 52; i++ {
		numbers = append(numbers, i)
	}

	for _, v := range numbers {
		fmt.Printf("Value: %v\t Type: %T\n", v, v)
	}

	slice1 := numbers[:5]
	slice2 := numbers[5:]
	slice3 := numbers[2:7]
	slice4 := numbers[1:6]

	fmt.Printf("%v\n", slice1)
	fmt.Printf("%v\n", slice2)
	fmt.Printf("%v\n", slice3)
	fmt.Printf("%v\n", slice4)
}
