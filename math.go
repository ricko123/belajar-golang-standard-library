package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(1.40))  //bulatkan ke atas
	fmt.Println(math.Floor(1.60)) //bulatkan ke kebawah
	fmt.Println(math.Round(1.60)) //bulatkan ke atas dan ke bawah mencari yang terdekat
	fmt.Println(math.Max(10, 11)) //mencari yang terbesar
	fmt.Println(math.Min(10, 11)) //mencari yang terkecil

}
