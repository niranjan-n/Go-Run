package main

import "fmt"

func foo(s string, c chan string) {
	c <- "Hi " + s
}
func main() {

	ch := make(chan string)

	go foo("Niru", ch)
	go foo("Tom", ch)
	v1, v2 := <-ch, <-ch
	fmt.Println(v1)
	fmt.Println(v2)
}
