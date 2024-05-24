package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	fmt.Printf("Value and Type of variable a: %v\t%T\n", a, a)
	fmt.Printf("Value and Type of variable b: %v\t%T\n", b, b)
	fmt.Printf("Value and Type of variable c: %v\t%T\n", c, c)
	fmt.Printf("Value and Type of variable d: %v\t%T\n", d, d)

	fmt.Println(*a)
	fmt.Println(*b)
	fmt.Println(*c)
	fmt.Println(*d)
}
