package main

import "fmt"

func sayHello(){
	fmt.Println("Hello World")
}

// function dengan parameter
func sayHelloTo(firstName string, lastName string){
	fmt.Println("Hello", firstName, lastName)
}

// return value
func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

// multiple value
func getFullName() (string, string) {
	return "Hermawan", "Budianto"
}

// named return value
func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Hermawan"
	middleName = "Budianto"
	lastName = "Permadi"

	return firstName, middleName, lastName
}


func main(){
	sayHello()
	sayHelloTo("Hermawan", "Budianto")
	
	result := getHello("Semua")
	fmt.Println(result)

	// function value
	hello := getHello
	fmt.Println(hello("Budi"))

	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	// hanya ambil salah satu
	_, firstName2 := getFullName()
	fmt.Println(firstName2)

	firstName3, middleName3, lastName3 := getCompleteName()
	fmt.Println(firstName3, middleName3, lastName3)
}