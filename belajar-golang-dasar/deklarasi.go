package main

import "fmt"

func main(){
	type NoKTP string

	var ktpHermawan NoKTP = "12345678"
	fmt.Println(ktpHermawan)
	fmt.Println(NoKTP("87654321"))
}