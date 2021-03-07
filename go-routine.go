package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello("Hey")
	go sayHello("there!!")

	time.Sleep(time.Millisecond)
}

func sayHello(s string) {

	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond)
	}
}
