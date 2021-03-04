package main

import "fmt"

type Person struct {
	first_name string
	last_name  string
	age        int
}

func main() {

	person1 := Person{
		first_name: "Niranjan",
		last_name:  "Hegde",
		age:        23,
	}

	fmt.Println(person1)
	fmt.Println(person1.age)

	person2 := Person{"Tom", "Cruise", 50}
	fmt.Println(person2)

	person3 := Person{
		first_name: "Heisenberg",
		age:        45,
	}
	fmt.Println(person3)
	fmt.Println("Last name :", person3.last_name)
}
