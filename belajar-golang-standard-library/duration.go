package main

import (
	"fmt"
	"time"
)

func main() {
	duration1 := time.Second * 100
	duration2 := time.Millisecond * 10
	duration3 := duration1 - duration2

	fmt.Printf("duration1: %d\n", duration3)
}