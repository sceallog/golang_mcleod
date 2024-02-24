package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(251)

	if x >= 0 && x <= 100 {
		fmt.Println("x is between 0 and 100")
	}

	if x >= 101 && x <= 200 {
		fmt.Println("x is between 101 and 200")
	}

	if x >= 201 && x <= 250 {
		fmt.Println("x is between 201 and 250")
	}

	// y := rand.Intn(3)
	// fmt.Printf("y = %d\n", y)
}
