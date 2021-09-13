package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time")

	present := time.Now()
	fmt.Println(present)
	// date time day
	fmt.Println(present.Format("15:04:05"))

	createdDate := time.Date(2020, time.July, 10, 23, 22, 0, 0, time.UTC)
	fmt.Println(createdDate)
}
