package main

import (
	"fmt"
)

func foo() {
	defer fmt.Println("Done!!")
	fmt.Println("Hey there!!")
	defer fmt.Println("are we done!!")

}

func fun() {

	for i := 0; i < 5; i++ {
		defer fmt.Println(i) //prints i in reverse order  i.e LIFO
	}
}
func main() {

	foo()

	fun()
}
