package list

import "github.com/nanafox/singly-linked-list/pkg/node"

// List is the structure for the singly linked list
type List struct {
	Head *node.Node // The head node of the list (denotes the beginning)
	Tail *node.Node // the tail node of the list (denotes the end)
	Size int        // Tracks the size of the list as it grows
}
