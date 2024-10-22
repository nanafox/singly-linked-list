package linkedList

import (
	"fmt"
)

// Print prints a visual representation of the linked list.
func (list List) Print() {
	if list.Head == nil || list.Size == 0 {
		fmt.Println("nil")
		return
	}

	currentNode := list.Head

	for currentNode != nil {
		fmt.Printf("%s ~> ", currentNode.nodeInfo())

		currentNode = currentNode.Next
	}

	fmt.Println("nil")
}

// PrintNodeInfo prints information about a specific node in the list.
func (node Node) PrintNodeInfo() {
	fmt.Printf("%s\n", node.nodeInfo())
}

// nodeInfo returns a string representation of the current node.
func (node Node) nodeInfo() string {
	return fmt.Sprintf("%s (%d, %s)", node.Name, node.Age, node.Passion)
}
