package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello,", name, "my name is", customer.Name)
}

func main(){
	hermawan := Customer{Name: "Hermawan"}
	hermawan.sayHello("Budi")
}