package main

import "fmt"

func main() {
  a, b, c := 747, 911, 90210

  fmt.Printf("%v \t\t in decimal: %d; \t in binary: %b; \t in hexadecimal: %#X; \n", a, a, a, a)
  fmt.Printf("%v \t\t in decimal: %d; \t in binary: %b; \t in hexadecimal: %#X; \n", b, b, b, b)
  fmt.Printf("%v \t\t in decimal: %d; \t in binary: %b; \t in hexadecimal: %#X; \n", c, c, c, c)
}
