package main

import "fmt"

type person struct {
	firstName      string
	lastName       string
	favIceFlavours []string
}

func main() {
	p1 := person{
		firstName:      "James",
		lastName:       "Bond",
		favIceFlavours: []string{"chocolate", "vanilla", "strawberry"},
	}

	p2 := person{
		firstName:      "Miss",
		lastName:       "Moneypenny",
		favIceFlavours: []string{"lemon", "mint", "hazelnut"},
	}

	fmt.Printf("%s %s's favourite ice cream flavours are:\n", p1.firstName, p1.lastName)
	for i, flav := range p1.favIceFlavours {
		fmt.Printf("  %d. %s\n", i+1, flav)
	}

	fmt.Printf("%s %s's favourite ice cream flavours are:\n", p2.firstName, p2.lastName)
	for i, flav := range p2.favIceFlavours {
		fmt.Printf("  %d. %s\n", i+1, flav)
	}
}
