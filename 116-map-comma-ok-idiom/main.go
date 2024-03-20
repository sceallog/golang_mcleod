package main

import "fmt"

func main() {
	am := map[string]int{
		"Todd":   42,
		"Henry":  16,
		"Padget": 14,
	}

	fmt.Println("--- accessing keys that don't exist ---")
	fmt.Println(am["George"])

	fmt.Println("--- using the comma, ok idiom ---")

	/* Could use comma ok idiom to check if the key exists like below

	v, ok := am["Henry"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key doesn't exist")
	}

	*/

	// However, it is more idiomatic Go syntax to use a statement statement idiom to limit the scope of the variables:
	if v, ok := am["George"]; !ok {
		fmt.Println("Key doesn't exist")
	} else {
		fmt.Println(v)
	}
}
