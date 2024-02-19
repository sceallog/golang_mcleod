package main

import (
	"fmt"

	"github.com/GoesToEleven/puppy"
)

func main() {
	var a int64 = 42
	var b int64 = 13
	var c int64 = 79

	fmt.Println(a + b + c)

	fmt.Println(puppy.Bark())
}
