package main

import "fmt"

func main() {
	fmt.Println(add(4, 6))
}

var add = func(a, b int) int {
	return a + b
}
