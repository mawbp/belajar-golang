package main

import "fmt"

func main(){
	var a = 10
	var b = 10
	var c = 2
	var x = a + b / c
	fmt.Println(x)

	var y = 10
	y += 30
	fmt.Println(y);

	var z = 5
	z++
	z++
	fmt.Println(z)
}