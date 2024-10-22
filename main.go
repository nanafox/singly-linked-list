package main

import (
	"fmt"
	"singlyLinkedList/linkedList"
)

func main() {
	list := linkedList.NewList()

	linkedList.PrintLinkedList(list)

	linkedList.Append(list, "Maxwell", 12, "Change the world")
	linkedList.Append(list, "Batman", 32, "Clean up Gotham of crime")

	linkedList.Prepend(list, "Clark Kent", 43, "Make the world a better place")

	linkedList.PrintLinkedList(list)

	node := linkedList.AtIndex(list, 1)

	linkedList.PrintNodeInfo(node)

	linkedList.DeleteAt(list, 5)
	linkedList.PrintLinkedList(list)

	linkedList.PrintReverse(list.Head)
	fmt.Println("nil")

	linkedList.InsertAt(list, 2, "Diana Prince", 1200, "Make man's world a better place")
	linkedList.PrintLinkedList(list)
}
