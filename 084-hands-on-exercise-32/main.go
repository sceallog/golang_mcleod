package main

import "fmt"

func main() {
	x := 35
	for {
		fmt.Println(x)
		if x < 20 {
			break
		}
		x--
	}
}
