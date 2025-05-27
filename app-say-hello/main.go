package main

import (
	"fmt"
	hello "github.com/mawbp/go-say-hello/v2"
)

func main() {
	fmt.Println(hello.SayHello("hermawan"))
	fmt.Println(hello.SayHelloWorld())
}