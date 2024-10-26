package list

import (
	"errors"
	"fmt"
	"github.com/nanafox/singly-linked-list/pkg/node"
)

type UninitializedError struct {
	method string
}

func (e *UninitializedError) Error() string {
	return fmt.Sprintf("%s: list is uninitialized", e.method)
}

// New initializes a new list instance to be used. It returns the pointer
// to the newly created list
func New() *List {
	return &List{nil, nil, 0}
}

// Append adds a new node with the given name, age, and passion to the
// end of the list. It updates the list's head and tail pointers and
// increments the list size.
func (list *List) Append(name string, age int, passion string) (
	*node.Node, error,
) {
	if list == nil {
		return nil, &UninitializedError{method: "list.Append"}
	}
	newNode := node.Create(name, age, passion)

	if list.Head == nil {
		list.Head = newNode
	} else {
		list.Tail.Next = newNode
	}

	list.Size += 1
	list.Tail = newNode

	return newNode, nil
}

// Prepend adds a new node at the beginning of the list.
func (list *List) Prepend(name string, age int, passion string) (
	*node.Node, error,
) {
	if list == nil {
		return nil, &UninitializedError{method: "list.Prepend"}
	}

	newNode := node.Create(name, age, passion)

	if list.Head == nil {
		list.Tail = newNode
	} else {
		newNode.Next = list.Head
	}

	list.Head = newNode
	list.Size += 1

	return newNode, nil
}

// AtIndex Returns the node at the given index
func (list *List) AtIndex(index int) (*node.Node, error) {
	if list.Head == nil || index >= list.Size {
		return nil, errors.New("invalid index")
	}

	if index == 0 {
		return list.Head, nil
	}

	if index == list.Size-1 {
		return list.Tail, nil
	}

	currentIndex := 0
	currentNode := list.Head

	for currentIndex < index {
		currentIndex += 1
		currentNode = currentNode.Next
	}

	return currentNode, nil
}

// DeleteAt deletes the node or element at the given index.
// Nothing is removed if the index provided is out of bounds, and an error
// message is printed.
//
// On a successful node removal, no error is returned. Errors are returned when
// the operation fails.
func (list *List) DeleteAt(index int) error {
	if list.Size == 0 {
		return errors.New("list.DeleteAt: cannot delete from an empty list")
	}

	if index == 0 {
		list.Head = list.Head.Next
		list.Size -= 1

		list.isEmpty()

		return nil
	}

	nodeToDelete, err := list.AtIndex(index)

	if err != nil {
		return err
	}

	if list.Size > 1 {
		nodeBefore, _ := list.AtIndex(index - 1)

		nodeBefore.Next = nodeToDelete.Next
	}

	list.Size -= 1

	list.isEmpty()
	return nil
}

// reversePrintHelper is helper function to print the nodes in reverse

// InsertAt inserts a node with the name, age, and passion at the given position.
// It returns the newly inserted node and no errors if the insertion is
// successful. On error, the node portion is nil and an error is returned.
//
// Note: Negative indexing is not supported (yet).
func (list *List) InsertAt(
	index int, name string, age int, passion string,
) (*node.Node, error) {
	if index < 0 {
		return nil, errors.New("InsertAt: negative indexing not supported")
	}

	if index == 0 {
		return list.Prepend(name, age, passion)
	}

	if index >= list.Size {
		return list.Append(name, age, passion)
	}

	newNode := node.Create(name, age, passion)
	nodeBefore, _ := list.AtIndex(index - 1)

	newNode.Next = nodeBefore.Next
	nodeBefore.Next = newNode

	list.Size += 1

	return newNode, nil
}

// Print prints a visual representation of the linked list.
func (list *List) Print() {
	if list.Head == nil || list.Size == 0 {
		fmt.Println("nil")
		return
	}

	currentNode := list.Head

	for currentNode != nil {
		fmt.Printf("%s ~> ", currentNode.Info())

		currentNode = currentNode.Next
	}

	fmt.Println("nil")
}

// PrintReverse prints the list in reverse without tampering with the original
// order of the list.
func (list *List) PrintReverse() {
	if list.Head == nil {
		return
	}

	node.ReversePrintHelper(list.Head)
}

// isEmpty checks to verify if the list empty or not. Once it verifies the size
// of the list is zero, it means the list is empty. It then sets the head and
// tail pointers to nil and return `true`. `false` is returned otherwise.
func (list *List) isEmpty() bool {
	if list.Size == 0 {
		list.Head = nil
		list.Tail = nil

		return true
	}

	return false
}
