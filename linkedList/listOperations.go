package linkedList

import (
	"fmt"
)

// NewList initializes a new list instance to be used. It returns the pointer
// to the newly created list
func NewList() *List {
	return &List{nil, nil, 0}
}

// createNode creates a node to be used by another functions
func createNode(name string, age int, passion string) Node {
	return Node{
		Name:    name,
		Age:     age,
		Passion: passion,
	}
}

// Append adds a new node with the given name, age, and passion to the
// end of the list. It updates the list's head and tail pointers and
// increments the list size.
func Append(list *List, name string, age int, passion string) bool {
	node := createNode(name, age, passion)

	if list.Head == nil {
		list.Head = &node
	} else {
		list.Tail.Next = &node
	}

	list.Size += 1
	list.Tail = &node

	return true
}

// Prepend adds a new node at the beginning of the list.
func Prepend(list *List, name string, age int, passion string) bool {
	node := createNode(name, age, passion)

	if list.Head == nil {
		list.Tail = &node
	} else {
		node.Next = list.Head
	}

	list.Head = &node
	list.Size += 1

	return true
}

// AtIndex Returns the node at the given index
func AtIndex(list *List, index int) *Node {
	if list.Head == nil || index >= list.Size {
		return nil
	}

	if index == 0 {
		return list.Head
	}

	if index == list.Size-1 {
		return list.Tail
	}

	currentIndex := 0
	currentNode := list.Head

	for currentIndex < index {
		currentIndex += 1
		currentNode = currentNode.Next
	}

	return currentNode
}

// DeleteAt deletes the node or element at the given index.
// Nothing is removed if the index provided is out of bounds, and an error
// message is printed.
//
// On a successful node removal, `true` is returned to signal the element was
// removed. `false` is returned if any error occurred.
func DeleteAt(list *List, index int) bool {
	if index == 0 {
		list.Head = list.Head.Next
		list.Size -= 1
		return true
	}

	nodeToDelete := AtIndex(list, index)

	if nodeToDelete == nil {
		fmt.Println("Element not found, nothing deleted")
		return false
	}

	if list.Size > 1 {
		nodeBefore := AtIndex(list, index-1)

		nodeBefore.Next = nodeToDelete.Next
	}

	list.Size -= 1
	return true
}

// PrintReverse prints the list in reverse without tampering with the original
// order of the list.
func PrintReverse(headNode *Node) {
	if headNode == nil {
		return
	}

	PrintReverse(headNode.Next)
	fmt.Printf("%s ~> ", nodeInfo(headNode))
}

// InsertAt inserts a node with the name, age, and passion at the given position.
// This function will return `true` on a successful insertion, `false` otherwise.
//
// Note: Negative indexing is not supported (yet).
func InsertAt(
	list *List, index int, name string, age int, passion string,
) bool {
	if index < 0 {
		fmt.Println("Negative indexing not supported")
		return false
	}

	if index == 0 {
		return Prepend(list, name, age, passion)
	}

	if index >= list.Size {
		return Append(list, name, age, passion)
	}

	newNode := createNode(name, age, passion)
	nodeBefore := AtIndex(list, index-1)

	newNode.Next = nodeBefore.Next
	nodeBefore.Next = &newNode

	list.Size += 1

	return true
}
