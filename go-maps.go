package main

import "fmt"

func main() {

	Footballers := make(map[string]string)
	Footballers["Striker"] = "Mbappe"
	Footballers["LWF"] = "Cr7"
	Footballers["RWF1"] = "Messi"
	Footballers["RWF2"] = "Salah"
	fmt.Println(Footballers)

	delete(Footballers, "RWF2")
	fmt.Println(Footballers)
	Footballers["RWF2"] = "Salah"

	for k, v := range Footballers {
		fmt.Println(k, ":", v)
		if k == "RWF2" {
			Footballers["LWF2"] = "Neymar"
		}
		if v == "Messi" {
			Footballers["GK"] = "Buffon"
		}

	}
	fmt.Println("Footballers: ", Footballers)

}
