package linkedList

import (
	"fmt"
)

// PrintLinkedList Prints a visual representation of the linked list
func PrintLinkedList(list *List) {
	if list == nil || list.Size == 0 {
		fmt.Println("nil")
		return
	}

	currentNode := list.Head

	for currentNode != nil {
		fmt.Printf(
			"%s (%d, %s) ~> ",
			currentNode.Name, currentNode.Age, currentNode.Passion,
		)

		currentNode = currentNode.Next
	}

	fmt.Println("nil")
}

// PrintNodeInfo Prints information about a specific node in the list
func PrintNodeInfo(node *Node) {
	fmt.Printf(
		"%s (%d, %s)\n",
		node.Name, node.Age, node.Passion,
	)
}
