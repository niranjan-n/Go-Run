package main

import (
	"fmt"
	"math"
	"math/rand"
)

func sqrt(s float64) {

	fmt.Println("Sqrt of s - ", math.Sqrt(s))
}

func genRandomNumber() {
	fmt.Println(rand.Intn(100))
}

func main() {
	sqrt(4)
	genRandomNumber()
}
