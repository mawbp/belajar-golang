package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Hermawan Budianto", "Hermawan"))
	fmt.Println(strings.Split("Hermawan Budianto", " "))
	fmt.Println(strings.ToLower("Hermawan Budianto"))
	fmt.Println(strings.ToUpper("Hermawan Budianto"))
	fmt.Println(strings.Trim("         Hermawan Budianto          ", " "))
	fmt.Println(strings.ReplaceAll("Hermawan Satu Hermawan Dua Hermawan Tiga", "Hermawan", "Budi"))
}