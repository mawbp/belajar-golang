package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Freddie", "Bonnie", "Chica", "Foxxy"}
	values := []int{100, 97, 85, 99}

	fmt.Println(slices.Min(values))
	fmt.Println(slices.Min(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Contains(names, "Chica"))
	fmt.Println(slices.Index(names, "William"))
	fmt.Println(slices.Index(names, "Foxxy"))
}
