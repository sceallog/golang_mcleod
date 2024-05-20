package main

import "fmt"

func main() {
	st := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "John",
		friends: map[string]int{
			"Mary":  25,
			"Peter": 28,
		},
		favDrinks: []string{"coffee", "tea", "water"},
	}

	fmt.Println(st)
	fmt.Printf("%#v\n", st)
	fmt.Println(st.first, st.friends["Mary"], len(st.friends), st.favDrinks[2])
}
