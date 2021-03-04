package main

import "fmt"

func main() {

	x := 10
	a := &x //holds the address

	fmt.Println(a)
	fmt.Println(*a)
	*a = 5
	fmt.Println(x)
	*a = *a * *a
	fmt.Println(x)
	fmt.Println(*a)
}
