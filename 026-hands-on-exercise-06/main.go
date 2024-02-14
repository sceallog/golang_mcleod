package main

import "fmt"

func main() {
  a := "First value"
  b := 123
  c := 5.75

  fmt.Printf("%v is of type %T \n", a, a)
  fmt.Printf("%v is of type %T \n", b, b)
  fmt.Printf("%v is of type %T \n", c, c)
}
