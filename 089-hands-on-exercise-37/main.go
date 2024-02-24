package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for k, v := range m {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}

	fmt.Println("--------------------")
	age := m["James"]
	fmt.Println(age)

	age2 := m["Q"]
	fmt.Println(age2)

	age3, ok := m["Q"]
	if ok {
		fmt.Println(age3)
	} else {
		fmt.Println("Q is not in the map")
	}
}
