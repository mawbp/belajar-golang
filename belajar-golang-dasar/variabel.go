package main

import "fmt"

func main(){
	var (
		firstName = "Ini First Name"
		lastName = "Ini Last Name"
	)

	// hanya pertama inisialisasi pake :=
	name := "Hermawan"
	umur := "20 Tahun"
	prodi := "Sistem informasi"

	fmt.Println(name)
	fmt.Println(umur)
	fmt.Println(prodi)
	fmt.Println("=======================")

	name = "Budianto"
	umur = "22 Tahun"
	prodi = "Matematika"

	fmt.Println(name)
	fmt.Println(umur)
	fmt.Println(prodi)

	fmt.Println(firstName)
	fmt.Println(lastName)
}