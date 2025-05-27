package main

import "fmt"

func main(){
	// isi dalam array di golang tidak dapat dihapus
	// hanya dirubah nilainya menjadi kosong
	var names [3]string
	names[0] = "Hermawan"
	names[1] = "Budianto"
	names[2] = "Permadi"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// tidak dibatasi panjang array nya
	var values = [...]int {90, 89, 70}

	fmt.Println(values)
	fmt.Println(len(values))
	values[0] = 100
	fmt.Println(values)
}