package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Printf("My name is %s and I'm on a walk.\n", d.first)
}

func (d *dog) run() {
	d.first = "Rover"
	fmt.Printf("My name is %s and I'm running.\n", d.first)
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.run()
}

func main() {
	d1 := dog{"Henry"}
	d2 := &dog{"Padget"}

	d1.walk()
	d1.run()
	//youngRun(d1)

	d2.walk()
	d2.run()
	youngRun(d2)
}
