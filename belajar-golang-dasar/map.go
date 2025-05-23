package main

import "fmt"

func main(){
	person := map[string]string{
		"name": "Hermawan",
		"address": "Jakarta",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Hermawan"
	book["page"] = "200"

	delete(book, "page")
	fmt.Println(book)
}