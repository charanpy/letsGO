package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter any stuff")

	// comma ok || error ok

	input, error := reader.ReadString('\n')
	fmt.Println("Stuff, ", input, error)
}
