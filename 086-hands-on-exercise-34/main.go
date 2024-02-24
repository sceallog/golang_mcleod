package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("Outer Loop Iteration %d \n", i)
		for j := 0; j < 5; j++ {
			fmt.Printf("\t Inner Loop Iteration %d \n", j)
		}
		fmt.Println("")
	}
}
