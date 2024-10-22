package main

import (
	"fmt"
	"singlyLinkedList/linkedList"
)

func main() {
	list := linkedList.NewList()

	list.Print()

	list.Append("Maxwell", 12, "Change the world")
	list.Append("Batman", 32, "Clean up Gotham of crime")

	list.Prepend("Clark Kent", 43, "Make the world a better place")

	list.Print()

	node := list.AtIndex(1)

	node.PrintNodeInfo()

	list.DeleteAt(1)
	list.Print()

	list.PrintReverse()
	fmt.Println("nil")

	list.InsertAt(2, "Diana Prince", 1200, "Make man's world a better place")
	list.Print()
}
