package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func sayHello(s string) {

	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond)
	}
	waitGroup.Done()
}

func main() {
	waitGroup.Add(1)
	go sayHello("Hey")
	waitGroup.Add(1)
	go sayHello("there!!")
	waitGroup.Wait()
	waitGroup.Add(1)
	go sayHello("Buddy")
	waitGroup.Wait()

}
