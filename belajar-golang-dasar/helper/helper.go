package helper

var version = "1.0.0" // tidak bisa diakses dari package lain
var Application = "golang"

// tidak bisa diakses dari luar package
func sayGoodBye(name string) string {
	return "Bye " + name
}

func SayHello(name string) string {
	return "hello " + name
}