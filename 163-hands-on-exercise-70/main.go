package main

import "fmt"

func main() {
	a := meaningOfLife()
	fmt.Println(a())
}

func meaningOfLife() func() int {
	return func() int {
		return 42
	}
}
