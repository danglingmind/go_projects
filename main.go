package main

import (
	"./practice"
)

func main() {
	// linked list ops
	items := make([]int, 0)
	for i := 1; i <= 10; i++ {
		items = append(items, i)
	}

	head := practice.CreateList(items)
	// print the list
	// practice.PrintList(head)
	// reverse the list
	// reversedList := practice.Reverse(head)
	// print the list
	// practice.PrintList(reversedList)
	mnrev := practice.ReverseMN(head, 2, 5)
	practice.PrintList(mnrev)

}
