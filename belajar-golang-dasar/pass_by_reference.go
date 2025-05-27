package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Surabaya", "Jawa Timur", "Indonesia"}
	address2 := &address1 // address2 pointer ke address1

	address2.City = "Bandung"

	fmt.Println(address1) // address1 ikut berubah berubah
	fmt.Println(address2)
}