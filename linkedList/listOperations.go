package linkedList

import (
	"fmt"
)

// NewList initializes a new list instance to be used. It returns the pointer
// to the newly created list
func NewList() *List {
	return &List{nil, nil, 0}
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

// Append adds a new node with the given name, age, and passion to the
// end of the list. It updates the list's head and tail pointers and
// increments the list size.
func (list *List) Append(name string, age int, passion string) bool {
	node := createNode(name, age, passion)

	if list.Head == nil {
		list.Head = node
	} else {
		list.Tail.Next = node
	}

	list.Size += 1
	list.Tail = node

	return true
}

// Prepend adds a new node at the beginning of the list.
func (list *List) Prepend(name string, age int, passion string) bool {
	node := createNode(name, age, passion)

	if list.Head == nil {
		list.Tail = node
	} else {
		node.Next = list.Head
	}

	list.Head = node
	list.Size += 1

	return true
}

// AtIndex Returns the node at the given index
func (list List) AtIndex(index int) *Node {
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
func (list *List) DeleteAt(index int) bool {
	if index == 0 {
		list.Head = list.Head.Next
		list.Size -= 1
		return true
	}

	nodeToDelete := list.AtIndex(index)

	if nodeToDelete == nil {
		fmt.Println("Element not found, nothing deleted")
		return false
	}

	if list.Size > 1 {
		nodeBefore := list.AtIndex(index - 1)

		nodeBefore.Next = nodeToDelete.Next
	}

	list.Size -= 1
	return true
}

// PrintReverse prints the list in reverse without tampering with the original
// order of the list.
func (list List) PrintReverse() {
	if list.Head == nil {
		return
	}

	reversePrintHelper(list.Head)
}

// reversePrintHelper is helper function to print the nodes in reverse
func reversePrintHelper(headNode *Node) {
	if headNode == nil {
		return
	}

	reversePrintHelper(headNode.Next)
	fmt.Printf("%s ~> ", headNode.nodeInfo())
}

// InsertAt inserts a node with the name, age, and passion at the given position.
// This function will return `true` on a successful insertion, `false` otherwise.
//
// Note: Negative indexing is not supported (yet).
func (list *List) InsertAt(
	index int, name string, age int, passion string,
) bool {
	if index < 0 {
		fmt.Println("Negative indexing not supported")
		return false
	}

	if index == 0 {
		return list.Prepend(name, age, passion)
	}

	if index >= list.Size {
		return list.Append(name, age, passion)
	}

	newNode := createNode(name, age, passion)
	nodeBefore := list.AtIndex(index - 1)

	newNode.Next = nodeBefore.Next
	nodeBefore.Next = newNode

	list.Size += 1

	return true
}
