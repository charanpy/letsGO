package main

import "fmt"

func main() {
	var username string = "charan"
	fmt.Println(username)
	fmt.Printf("Vaiable is of type: %T \n", username)

	var isLoggdIn bool = false
	fmt.Printf("%T\n", isLoggdIn)

	var smallVal uint8 = 255
	fmt.Printf("%T\n", smallVal)

	var smallFloat float32 = 255.45666
	fmt.Printf("%T", smallFloat)

	// default values and aliases
	var anotherVar int
	fmt.Printf("%T %d\n", anotherVar, anotherVar)

	var anotherStr string
	fmt.Printf("1 %s", anotherStr)
	fmt.Printf("%T", anotherStr)

	// implicit
	var website = "https"
	fmt.Println(website)

	// no var style
	numberOfUser := 300
	fmt.Printf(numberOfUser)

}
