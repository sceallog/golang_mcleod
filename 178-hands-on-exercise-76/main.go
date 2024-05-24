package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Printf("My name is %s and I'm on a walk.\n", d.first)
}

func (d dog) run() {
	fmt.Printf("My name is %s and I'm running.\n", d.first)
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.walk()
	y.run()
}

func (d dog) changeName(s string) dog {
	d.first = s
	return d
}

func main() {
	d1 := dog{"Henry"}
	youngRun(d1)

	d2 := dog{"Padget"}
	youngRun(d2)
	d2 = d2.changeName("Rover")
	youngRun(d2)
}
