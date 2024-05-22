package main

import "fmt"

// Value semantics
func addOne(v int) int {
	return v + 1
}

// Pointer semantics
func addOneP(v *int) {
	*v += 1
}

func main() {
	// Value semantics
	a := 1
	fmt.Println(a)         // 1
	fmt.Println(addOne(a)) // 2
	fmt.Println(a)         //1

	// Pointer semantics
	b := 1
	fmt.Println(b)
	addOneP(&b)
	fmt.Println(b)
}
