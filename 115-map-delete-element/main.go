package main

import "fmt"

func main() {
	am := map[string]int{
		"Todd":   42,
		"Henry":  16,
		"Padget": 14,
	}

	fmt.Println(am)

	delete(am, "Todd")

	for k, v := range am {
		fmt.Println(k, v)
	}

	for _, v := range am {
		fmt.Println(v)
	}

	for k := range am {
		fmt.Println(k)
	}
}
