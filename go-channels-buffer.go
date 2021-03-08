package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo(i int, c chan int) {
	defer wg.Done()
	c <- i * 2

}
func main() {

	ch := make(chan int, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(i, ch)
	}
	wg.Wait()
	close(ch)

	for value := range ch {
		fmt.Println(value)
	}
}
