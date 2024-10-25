package main

import (
	"fmt"
	"github.com/nanafox/singly-linked-list/pkg/list"
)

func main() {
	myList := list.New()

	_, err := myList.Append("Maxwell", 12, "Change the world")
	if err != nil {
		fmt.Println(err)
	}
	_, err = myList.Append("Batman", 32, "Clean up Gotham of crime")
	if err != nil {
		fmt.Println(err)
	}

	_, err = myList.Prepend("Clark Kent", 43, "Make the world a better place")
	if err != nil {
		fmt.Println(err)
	}

	myList.Print()

	node, _ := myList.AtIndex(1)

	node.PrintNodeInfo()

	err = myList.DeleteAt(1)
	if err != nil {
		fmt.Println(err)
	}

	myList.Print()

	myList.PrintReverse()
	fmt.Println("nil")

	node, err = myList.InsertAt(
		2, "Diana Prince", 1200, "Make man's world a better place",
	)
	if err != nil {
		fmt.Println(err)
	} else {
		node.PrintNodeInfo()
	}

	myList.Print()

	// Insertion using negative index returns an error
	node, err = myList.InsertAt(-5, "Nothin", 3, "blah")
	if err != nil {
		fmt.Println(err)
	}

	// Set myList to nil and try insertion
	myList = nil

	node, err = myList.Append("Hello", 12, "anything for the world")

	// this should print the error
	if err != nil {
		fmt.Println(err)
	}
}
