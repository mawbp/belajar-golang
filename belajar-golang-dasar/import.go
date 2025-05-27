package main

import (
	"belajar-golang-dasar/helper"
	"belajar-golang-dasar/database"
	_ "belajar-golang-dasar/internal" // hanya menjalankan init, tanpa menggunakan function lain
	"fmt"
)

func main() {
	result := helper.SayHello("Hermawan")
	fmt.Println(result)

	fmt.Println(helper.Application)
	fmt.Println(database.GetDatabase())

	/* tidak bisa diakses
	fmt.Println(helper.version)
	fmt.Println(helper.sayGoodBye("Hermawan"))
	*/
}