package main

import "fmt"

func main() {
	func(a, b int) {
		fmt.Println(a + b)
	}(5, 8)
}
