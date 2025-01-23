package main

import (
	"fmt"
	"slices"
)

func main() {
	name := []string{"joko", "Paul", "Chagatai", "Zengi"}
	values := []int{100, 90, 80, 40}

	fmt.Println(slices.Min(name))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(name))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(name, "Ricko"))
	fmt.Println(slices.Index(name, "Ricko"))
	fmt.Println(slices.Index(name, "Paul"))
}
