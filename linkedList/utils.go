package linkedList

import (
	"fmt"
)

// Print prints a visual representation of the linked list.
func (list *List) Print() {
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

// createNode creates a node to be used by another functions. It returns the
// address of the newly created node.
func createNode(name string, age int, passion string) *Node {
	return &Node{
		Name:    name,
		Age:     age,
		Passion: passion,
	}
}

// PrintReverse prints the list in reverse without tampering with the original
// order of the list.
func (list *List) PrintReverse() {
	if list.Head == nil {
		return
	}

	reversePrintHelper(list.Head)
}

// reversePrintHelper is a recursive function that prints the linked
// list in reverse order.
// It takes the head node of the list as an argument.
func reversePrintHelper(headNode *Node) {
	if headNode == nil {
		return
	}

	reversePrintHelper(headNode.Next)
	fmt.Printf("%s ~> ", headNode.nodeInfo())
}
