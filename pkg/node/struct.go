package node

// Node defines the node structure for the linked list data structure
type Node struct {
	Name    string // The name of the person
	Age     int    // The age of the person
	Passion string // The person's innermost desire and passion
	Next    *Node  // A pointer to the next node
}
