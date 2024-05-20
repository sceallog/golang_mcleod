package main

import "fmt"

func main() {
	foo()

	bar("Todd")

	s := baz("Todd")
	fmt.Println(s)

	fmt.Print(dogYears("Toodles", 12))
}

/*
function syntax

no params, no returns
1 param, no returns
1 param, 1 return
2 params, 2 returns
*/

func foo() {
	fmt.Println("I am from foo")
}

func bar(s string) {
	fmt.Println("My name is", s)
}

func baz(s string) string {
	return fmt.Sprint("They call me Mr ", s)
}

func dogYears(name string, ageDogYears int) (string, int) {
	ageDogYears *= 7
	return fmt.Sprint(name, " is this old in dog years: "), ageDogYears
}

// func (r receiver) identifier(p parameter(s)) (return(s)) {<code here>}
