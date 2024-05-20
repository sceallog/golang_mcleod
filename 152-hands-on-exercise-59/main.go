package main

import "fmt"

func main() {
	xi := []int{23, 24, 25, 26, 27, 28, 29}

	x := foo(23)
	fmt.Println(x)

	y := foo(xi...)
	fmt.Println(y)

	z := bar(xi)
	fmt.Println(z)
}

func foo(i ...int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}

func bar(ii []int) int {
	sum := 0
	for _, v := range ii {
		sum += v
	}
	return sum
}

/*
create a func with the identifier foo that
takes in a variadic parameter of type int
pass in a value of type []int into your func (unfurl the []int)
returns the sum of all values of type int passed in
create a func with the identifier bar that
takes in a parameter of type []int
returns the sum of all values of type int passed in
*/
