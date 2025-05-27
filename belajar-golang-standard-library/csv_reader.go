package main

import (
	"fmt"
	"encoding/csv"
	"strings"
	"io"
)

func main() {
	csvString := "hermawan,budianto,permadi\n" +
		"budi, pratama, dian\n" +
		"joko, moro, zeka"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}