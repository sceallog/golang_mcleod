package main

/*
● create a type SQUARE
	○ length float64
	○ width float64
● create a type CIRCLE
	○ radius float64
● attach a method to each that calculates AREA and returns it
	○ circle area= π r 2
		■ math.Pi
		■ math.Pow
	○ square area=L*W
● create a type SHAPE that defines an interface as anything that has the AREA method
● create a func INFO which takes type shape and then prints the area
● create a value of type square
● create a value of type circle
● use func info to print the area of square
● use func info to print the area of circle
*/
import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	sq := square{
		length: 10,
		width:  15,
	}

	circ := circle{
		radius: 4,
	}

	info(sq)
	info(circ)
}
