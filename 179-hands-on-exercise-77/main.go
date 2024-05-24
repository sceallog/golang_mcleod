package main

import "fmt"

type Person struct {
	first string
}

func changeFirstNameV(p Person, newName string) string {
	p.first = newName
	return p.first
}

func changeFirstNameP(p *Person, newName string) {
	p.first = newName
}

func main() {
	person1 := Person{first: "James"}
	fmt.Println(person1.first)

	person1.first = changeFirstNameV(person1, "Josh")
	fmt.Println(person1.first)

	changeFirstNameP(&person1, "James")
	fmt.Println(person1.first)
}
