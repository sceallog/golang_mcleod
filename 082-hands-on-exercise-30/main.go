package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 1; i < 43; i++ {
		x := rand.Intn(5)
		fmt.Printf("Iteration %d:\t", i)
		switch x {
		case 0:
			fmt.Println("x is 0")
		case 1:
			fmt.Println("x is 1")
		case 2:
			fmt.Println("x is 2")
		case 3:
			fmt.Println("x is 3")
		case 4:
			fmt.Println("x is 4")
		}
	}
}
