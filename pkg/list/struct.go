package list

import "github.com/nanafox/singly-linked-list/pkg/node"

// List is the structure for the singly linked list
type List struct {
	Head *node.Node
	Tail *node.Node
	Size int
}
