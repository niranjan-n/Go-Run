package main

import "fmt"

func arithmetic(a, b int) (int, int, int, int) {

	return a + b, a - b, a * b, a / b
}
func main() {

	num1, num2, num3, num4 := arithmetic(4, 2)
	fmt.Println(num1, num2, num3, num4)

	var s string = "Hi There!!!"
	fmt.Println(s)
}
