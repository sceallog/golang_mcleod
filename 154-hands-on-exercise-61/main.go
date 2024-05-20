package main

/*
● Create a user defined struct with
	○ the identifier “person”
	○ the fields:
		■ first
		■ age
● attach a method to type person with
	○ the identifier “speak”
	○ the method should have the person say their name and age
● create a value of type person
● call the method from the value of type person
*/
import "fmt"

type Person struct {
	first string
	age   int
}

func (p Person) speak() {
	fmt.Printf("My name is %s and I am %d years old.\n", p.first, p.age)
}

func main() {
	p1 := Person{
		first: "John",
		age:   33,
	}

	p1.speak()
}
