package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex = regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("oke"))
	fmt.Println(regex.MatchString("elo"))
	fmt.Println(regex.MatchString("eDo"))

	fmt.Println(regex.FindAllString("eko edo egi elu ego elo eto eka alo", 10))
}