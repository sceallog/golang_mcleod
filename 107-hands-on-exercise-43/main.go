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
}

/*
● Using a COMPOSITE LITERAL:
● create a SLICE of TYPE int
● assign these 10 VALUES
42, 43, 44, 45, 46, 47, 48, 49, 50, 51
● Range over the slice and print the values out.
○ print out the VALUE and the TYPE
*/
