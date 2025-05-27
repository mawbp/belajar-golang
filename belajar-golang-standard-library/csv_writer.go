package main

import (
	"encoding/csv"
	"os"
)

func main(){
	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"hermawan", "budianto", "permmadi"})
	_ = writer.Write([]string{"budi", "pratama", "dian"})
	_ = writer.Write([]string{"joko", "morro", "zeka"})
	writer.Flush()
}