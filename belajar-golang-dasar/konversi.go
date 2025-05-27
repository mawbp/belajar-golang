package main

import "fmt"

func main(){
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)

	// karena nilai melebihi kapasitas int16, maka akan mengulang nilai dari bawah
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)


	var name = "Hermawan"
	var h = name[0]
	var hString = string(h)

	fmt.Println(name)
	fmt.Println(hString)
}