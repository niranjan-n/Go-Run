package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func cleanUp() {
	if r := recover(); r != nil {
		fmt.Println("recovered from panic", r)
	}
	waitGroup.Done()

}
func sayHello(s string) {
	defer cleanUp()

	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond)
		if i == 3 {
			panic("Oh dear,it's a 2")
		}
	}
}

func main() {
	waitGroup.Add(1)
	go sayHello("Hey")
	waitGroup.Wait()

}
