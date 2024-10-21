package main

import (
	"singlyLinkedList/linkedList"
)

func main() {
	list := linkedList.NewList()

	linkedList.PrintLinkedList(list)

	linkedList.Append(list, "Maxwell", 12, "Change the world")
	linkedList.Append(list, "Batman", 32, "Clean up Gotham of crime")

	//linkedList.PrintLinkedList(&list)

	linkedList.Prepend(list, "Clark Kent", 43, "Make the world a better place")

	linkedList.PrintLinkedList(list)

	node := linkedList.AtIndex(list, 1)

	linkedList.PrintNodeInfo(node)

	linkedList.DeleteAt(list, 5)
	linkedList.PrintLinkedList(list)
}
