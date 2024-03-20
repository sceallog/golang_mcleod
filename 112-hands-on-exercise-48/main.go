package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mm := []string{"Miss", "Moneypenny", "I'm 008"}
	records := [][]string{jb, mm}

	for i, v := range records {
		fmt.Printf("%v - %v\n", i, v)
		for j, s := range v {
			fmt.Printf("%v - %v\n", j, s)
		}
	}
}
