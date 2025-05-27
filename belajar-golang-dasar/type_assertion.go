package main

import "fmt"

func random() interface{} {
	return "OK"
}

func main() {
	result := random()
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown")
	}

	/* panic
		resultInt := result.(int)
		fmt.Println(resultInt)
	*/
}