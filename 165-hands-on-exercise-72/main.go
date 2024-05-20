package main

import "fmt"

func main() {
	m := incrementalMultiplicator(3)
	fmt.Println(m())
	fmt.Println(m())
	fmt.Println(m())
	fmt.Println(m())
	fmt.Println(m())
	fmt.Println(m())
	fmt.Println(m())
	fmt.Println(m())
	fmt.Println(m())
	fmt.Println(m())
}

func incrementalMultiplicator(n int) func() int {
	multiplicator := 1
	return func() int {
		result := n * multiplicator
		multiplicator++
		return result
	}
}
