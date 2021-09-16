package main

import "fmt"

func main() {
	fmt.Println("Welcome")

	// var ptr *int

	// fmt.Println("Value of pointer is", ptr)

	myNumber := 23

	var ptr = &myNumber
	fmt.Println(ptr)
	fmt.Println(*ptr)
}
