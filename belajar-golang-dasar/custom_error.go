package main

import "fmt"

type validationError struct {
	Message string
}

type notFoundError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "validation error"}
	}

	if id != "hermawan" {
		return &notFoundError{Message: "data not found"}
	}

	return nil
}

func main() {
	err := SaveData("budi", nil)
	if err != nil {
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println("validation error:", validationErr.Message)
		// } else if notFoundErr, ok := err.(*notFoundError); ok {
		// 	fmt.Println("not found error:", notFoundErr.Message)
		// } else {
		// 	fmt.Println("unknown error:", err.Error())
		// }

		switch finalErr := err.(type) {
		case *validationError:
			fmt.Println("validation error:", finalErr.Error())
		case *notFoundError:
			fmt.Println("not found error:", finalErr.Error())
		default:
			fmt.Println("unknown error:", finalErr.Error())
		}
	} else {
		fmt.Println("Sukses")
	}
}