package main

import "fmt"

func main() {
	fmt.Println(printSquare(square, 14))
}

type callbackFunc func(int) int

func square(n int) int {
	return n * n
}

func printSquare(cb callbackFunc, n int) string {
	return fmt.Sprintf("%v", cb(n))
}
