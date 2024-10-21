package linkedList

import "fmt"

func NewList() *List {
	return &List{nil, nil, 0}
}

func createNode(name string, age int, passion string) Node {
	return Node{
		Name:    name,
		Age:     age,
		Passion: passion,
	}
}

func Append(list *List, name string, age int, passion string) {
	node := createNode(name, age, passion)

	if list.Head == nil {
		list.Head = &node
	} else {
		list.Tail.Next = &node
	}

	list.Size += 1
	list.Tail = &node
}

func Prepend(list *List, name string, age int, passion string) {
	node := createNode(name, age, passion)

	if list.Head == nil {
		list.Tail = &node
	} else {
		node.Next = list.Head
	}

	list.Head = &node
	list.Size += 1
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

func DeleteAt(list *List, index int) {
	if index == 0 {
		list.Head = list.Head.Next
		list.Size -= 1
		return
	}

	nodeToDelete := AtIndex(list, index)

	if nodeToDelete == nil {
		fmt.Println("Element not found, nothing deleted")
		return
	}

	if list.Size > 1 {
		nodeBefore := AtIndex(list, index-1)

		nodeBefore.Next = nodeToDelete.Next
	}

	list.Size -= 1
}
