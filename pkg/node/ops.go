package node

import (
	"fmt"
)

// Create returns a new node to be used by another functions. It returns the
// address of the newly created node.
func Create(name string, age int, passion string) *Node {
	return &Node{
		Name:    name,
		Age:     age,
		Passion: passion,
	}
}

// PrintInfo prints information about a specific node in the list.
func (node Node) PrintInfo() {
	fmt.Printf("%s\n", node.Info())
}

// Info returns a string representation of the current node.
func (node Node) Info() string {
	return fmt.Sprintf("%s (%d, %s)", node.Name, node.Age, node.Passion)
}

// ReversePrintHelper is a recursive function that prints the linked
// list in reverse order. It takes the head node of the list as an argument.
func ReversePrintHelper(headNode *Node) {
	if headNode == nil {
		return
	}

	ReversePrintHelper(headNode.Next)
	fmt.Printf("%s ~> ", headNode.Info())
}
