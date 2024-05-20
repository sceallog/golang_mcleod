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
		firstName:      "Jane",
		lastName:       "Moneypenny",
		favIceFlavours: []string{"lemon", "mint", "hazelnut"},
	}

	people := make(map[string]person)

	people[p1.lastName] = p1
	people[p2.lastName] = p2

	for k, v := range people {
		fmt.Printf("%s's full name and favourite ice cream flavours are:\n %s %s\n", k, v.firstName, v.lastName)
		for i, flav := range v.favIceFlavours {
			fmt.Printf("  %d. %s\n", i+1, flav)
		}
	}
}
