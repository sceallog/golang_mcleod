package main

import "fmt"

func main() {
	am := make(map[string][]string)

	am["bond_james"] = []string{"shaken, not stirred", "martinis", "fast cars"}
	am["moneypenny_miss"] = []string{"intelligence", "literature", "target practice"}
	am["no_dr"] = []string{"world domination", "opera", "cat stroking"}

	for k, v := range am {
		fmt.Printf("----------\n%s:\n", k)
		for i, el := range v {
			fmt.Printf("  %d - %s \n", i, el)
		}
	}
}
