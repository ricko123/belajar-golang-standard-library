package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("edo"))
	fmt.Println(regex.MatchString("Eko"))
	fmt.Println(regex.MatchString("eko"))

	fmt.Println(regex.FindAllString("eko, edo egi e10 eto eKo ego", 10))
}

// https://github.com/google/re2/wiki/syntax
