package main

import "fmt"

// meenggunakan for loop
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

// menggunakan recursive
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value - 1)
	}
}

func main(){
	result := 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1
	fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecursive(10))
	fmt.Println(result)
}