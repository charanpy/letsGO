package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
)

func main() {
	// random
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5))

	// crypto
	randNo, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(randNo)
}
