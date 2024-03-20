package main

import "fmt"

func main() {
	xs := []string{"hello", "world"}
	fmt.Println(xs)
}

/*
	A slice is a data structure with three elements:
	-- len
	-- cap
	-- pointer to an underlying array

	type slice struct {
		array unsafe.Pointer
		len   int
		cap 	int
	}
*/
