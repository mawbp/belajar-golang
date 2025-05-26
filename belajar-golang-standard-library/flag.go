package main

import (
	"fmt"
	"flag"
)

func main() {
	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "root", "Put your database password")
	port := flag.Int("port", 0, "Put your database port")

	flag.Parse()

	fmt.Println("Host", *host)
	fmt.Println("Username", *username)
	fmt.Println("Password", *password)
	fmt.Println("Port", *port)
}