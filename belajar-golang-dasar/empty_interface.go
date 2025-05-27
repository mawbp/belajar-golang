package main

import "fmt"

func Ups() interface{} {
	return "abc"
}

func main() {
	kosong := Ups()
	fmt.Println(kosong)
}