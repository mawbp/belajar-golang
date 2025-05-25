package main

import "fmt"

func main(){
	name := "hermawannnnn"

	switch name {
	case "Hermawan":
		fmt.Println("Hello Hermawan")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Hallo !!")
	}

	// switch dengan short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama lebih dari 5 huruf")
	case false:
		fmt.Println("Nama kurang dari / sama dengan 5 huruf")
	}

	//switch tanpa kondisi, kondisi ditulis tiap case
	panjang := len(name)
	switch {
	case panjang > 10:
		fmt.Println("Nama lebih dari 10 huruf")
	case panjang > 5:
		fmt.Println("Nama lebih dari 5 huruf")
	default:
		fmt.Println("Nama kurang dari / sama dengan 5 huruf")
	}
}