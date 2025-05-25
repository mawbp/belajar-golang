package main

import "fmt"

// function type declaration
type Filter func(string) string
type Blacklist func(string) bool

// function as parameter
func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello, ", filteredName)
}

func spanFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main(){
	filter := spanFilter
	sayHelloWithFilter("Anjing", filter)

	//anonymous function
	blacklist := func(name string) bool {
		return name == "anjing"
	}
	registerUser("hermawan", blacklist)
}