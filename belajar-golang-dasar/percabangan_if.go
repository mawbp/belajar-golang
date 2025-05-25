package main

import "fmt"

func main(){
	name := "Hermawan"
	if name == "Hermawan" {
		fmt.Println("Hello hermawan")
	} else if name == "Joko" {
		fmt.Println("Hello joko")
	} else {
		fmt.Println("Halo, !!")
	}

	// if dengan short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama lebih dari 5 huruf")
	} else {
		fmt.Println("Nama kurang dari 5 huruf")
	}
}