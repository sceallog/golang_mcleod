package main

import "fmt"

func main() {
	fmt.Println(Paradise("Okinawa"))
}

// Paradise takes a location as string and returns a string with the location in a sentence.
func Paradise(loc string) string {
	return fmt.Sprint("My idea of paradise is ", loc)
}
