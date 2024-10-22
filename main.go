package main

import (
	"fmt"
	"singlyLinkedList/linkedList"
)

func main() {
	list := linkedList.NewList()

	_, err := list.Append("Maxwell", 12, "Change the world")
	if err != nil {
		fmt.Println(err)
	}
	_, err = list.Append("Batman", 32, "Clean up Gotham of crime")
	if err != nil {
		fmt.Println(err)
	}

	_, err = list.Prepend("Clark Kent", 43, "Make the world a better place")
	if err != nil {
		fmt.Println(err)
	}

	list.Print()

	node, _ := list.AtIndex(1)

	node.PrintNodeInfo()

	err = list.DeleteAt(1)
	if err != nil {
		fmt.Println(err)
	}

	list.Print()

	list.PrintReverse()
	fmt.Println("nil")

	node, err = list.InsertAt(2, "Diana Prince", 1200, "Make man's world a better place")
	if err != nil {
		fmt.Println(err)
	} else {
		node.PrintNodeInfo()
	}

	list.Print()

	// Insertion using negative index returns an error
	node, err = list.InsertAt(-5, "Nothin", 3, "blah")
	if err != nil {
		fmt.Println(err)
	}

	// Set list to nil and try insertion
	list = nil

	node, err = list.Append("Hello", 12, "anything for the world")

	// this should print the error
	if err != nil {
		fmt.Println(err)
	}
}
