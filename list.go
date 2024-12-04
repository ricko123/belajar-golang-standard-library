package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Ricko")
	data.PushBack("Sapto")
	data.PushBack("Prihartanto")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) //Ricko

	next := head.Next() //Sapto
	fmt.Println(next.Value)

	next = next.Next() //Prihartanto
	fmt.Println(next.Value)
	//looping
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
