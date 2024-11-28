package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Ricko Sapto", "Sapto"))
	fmt.Println(strings.Split("Ricko Sapto", " "))
	fmt.Println(strings.ToLower("Ricko Sapto"))
	fmt.Println(strings.ToUpper("Ricko Sapto"))
	fmt.Println(strings.Trim("     Ricko Sapto    ", " "))
	fmt.Println(strings.ReplaceAll("Ricko Sapto Ricko", "Ricko", "Budi"))
}
