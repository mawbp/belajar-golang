package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func main() {
	var hermawan Customer
	hermawan.Name = "Hermawan"
	hermawan.Address = "Indonesia"
	hermawan.Age = 20
	fmt.Println(hermawan)

	joko := Customer {
		Name: "Joko",
		Address: "Indonesia",
		Age: 20,
	}
	fmt.Println(joko)

	budi := Customer {"Budi", "Indonesia", 30}
	fmt.Println(budi)
}