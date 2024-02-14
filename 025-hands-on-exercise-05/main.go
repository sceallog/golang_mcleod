package main

import "fmt"

func main() {
  var zero int
  short := "A string value"
  a, b, c := 1, "one", 1.0
  var negative int = -59

  fmt.Printf("%v \t\t %T \n", zero, zero)
  fmt.Printf("%v \t %T \n", short, short)
  fmt.Printf("%v \t %T %v \t %T %v \t %T\n", a, a, b, b, c, c)
  fmt.Printf("%v \t\t %T \n", negative, negative)
}
