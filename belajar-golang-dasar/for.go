package main

import "fmt"

func main(){

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	fmt.Println("Selesai")

	// for range
	names := []string{"Hermawan", "Budi", "Anto"}
	for index, name := range names {
		fmt.Println("Index:", index, "=", name)
	}

	//jika tidak butuh index
	for _, name := range names {
		fmt.Println(name)
	}


}