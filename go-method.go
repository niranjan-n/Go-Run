package main

import "fmt"

type Rectangle struct {
	length, breadth int
}

func (r Rectangle) area() int {
	return r.length * r.breadth
}

//value reciever method..Doesn't modify the length pf the rectangle
func (r Rectangle) doubleTheLength() int {
	return r.length * 2
}

//pointer reciever method..Doesn't modify the length pf the rectangle
func (r *Rectangle) doubleTheBreadth() {
	r.breadth *= 2
}

//function

func area(r Rectangle) int {
	return r.length * r.breadth
}

func main() {
	rec := Rectangle{20, 10}
	//calling method to calculate the area
	fmt.Println("area methodCall: ", rec.area())

	//calling function to calculate the area
	a := area(rec)
	fmt.Println("area functionCall: ", a)

	fmt.Println(rec.doubleTheLength())
	fmt.Println(rec) //  length remains unchanged

	rec.doubleTheBreadth()
	fmt.Println(rec) //  Breadth value  gets modified

}
