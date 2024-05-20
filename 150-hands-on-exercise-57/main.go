package main

import "fmt"

func main() {
	x := sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(x)
}

func sum(ii []int) int {
	t := 0
	for _, v := range ii {
		t += v
	}
	return t
}
