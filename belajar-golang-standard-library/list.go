package main

import (
	"fmt"
	"container/list" // implementasi struktur data double linked list
)

func main() {
	var data = list.New()
	data.PushBack("Hermawan")
	data.PushBack("Budianto")
	data.PushBack("Permadi")

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}