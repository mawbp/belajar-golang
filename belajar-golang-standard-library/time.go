package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Local())

	utc := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(utc.Local())

	parse, err := time.Parse(time.RFC3339, "2025-01-03T10:20:10Z")
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println(parse)
	}

	fmt.Println(parse.Year())
	fmt.Println(parse.Month())
	fmt.Println(parse.Day())
	fmt.Println(parse.Hour())

}