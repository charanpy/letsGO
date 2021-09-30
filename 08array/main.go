package main

import "fmt"

func main() {
	fmt.Println("welcome")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Peach"

	fmt.Println(fruitList, len(fruitList))

	var vegList = [3]string{"po", "to", "ve"}
	fmt.Println(vegList)

}
