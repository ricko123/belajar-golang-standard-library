package main

import "fmt"

func main() {
	firstName := "Ricko"
	lastName := "Prihartanto"

	fmt.Println("Hello '", firstName, lastName, "'")
	fmt.Printf("Hello '%s %s'\n", firstName, lastName)
}
