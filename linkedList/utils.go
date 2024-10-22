package linkedList

import (
	"fmt"
)

// PrintLinkedList prints a visual representation of the linked list.
func PrintLinkedList(list *List) {
	if list == nil || list.Size == 0 {
		fmt.Println("nil")
		return
	}

	currentNode := list.Head

	for currentNode != nil {
		fmt.Printf("%s ~> ", nodeInfo(currentNode))

		currentNode = currentNode.Next
	}

	fmt.Println("nil")
}

// PrintNodeInfo prints information about a specific node in the list.
func PrintNodeInfo(node *Node) {
	fmt.Printf("%s\n", nodeInfo(node))
}

// nodeInfo returns a string representation of the current node.
func nodeInfo(node *Node) string {
	return fmt.Sprintf("%s (%d %s)", node.Name, node.Age, node.Passion)
}
