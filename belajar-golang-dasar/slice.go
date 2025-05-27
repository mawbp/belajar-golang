package main

import "fmt"

func main(){
	names := [...]string{"Hermawan", "Budianto", "Permadi", "Budi", "Anto", "Dian"}
	slice := names[4:6]
	slice2 := names[:3]
	slice3 := names[3:]
	slice4 := names[:]

	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println("============================================")

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Upss"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)
	fmt.Println("============================================")

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Hermawan"
	newSlice[1] = "Hermawan"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(toSlice)
}