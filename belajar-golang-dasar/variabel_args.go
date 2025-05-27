package main

import "fmt"

// variadic function
func sumALl(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main(){
	total := sumALl(10, 10, 10, 10, 10, 10)
	fmt.Println(total)

	// jika data berupa slice
	numbers := []int{10, 10, 10}
	total2 := sumALl(numbers...)
	fmt.Println(total2)
}