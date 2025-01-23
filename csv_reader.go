package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "ricko, sapto, prihartanto\n" +
		"budi, pratama, nugraha\n" +
		"joko, morro, diah"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println((record))
	}
}
