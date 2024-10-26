package list

import (
	"errors"
	"testing"
)

func TestNew(t *testing.T) {
	newList := New()

	if newList.Head != nil {
		t.Fatal("The head of the list must be nil when initialized")
	}

	if newList.Tail != nil {
		t.Fatal("The tail of the list must be nil when initialized")
	}

	if newList.Size != 0 {
		t.Fatalf("The initial size of the list must be 0, not %d", newList.Size)
	}
}

func TestList_Append(t *testing.T) {
	t.Run(
		"Append 100 times", func(t *testing.T) {
			testList := New()
			want := 100 // Expected list size

			for i := 0; i < want; i++ {

				_, err := testList.Append("Person", 12, "Passion")
				if err != nil {
					t.Fatal("An error occurred during testing")
				}
			}

			got := testList.Size

			if got != want {
				t.Fatalf("Expected size of the list to %d, but got %d", want, got)
			}
		},
	)

	t.Run(
		"Append to a list of two nodes", func(t *testing.T) {
			testList := New()

			testList.Append("Person 1", 12, "Passion 1")
			testList.Append("Person 2", 16, "Passion 2")

			lastNode, _ := testList.Append("Person 3", 19, "Passion 3")

			want := "Person 3 (19, Passion 3)"
			got := lastNode.Info()

			if got != want {
				t.Fatalf("Expected the node info to be %s, but got %s", want, got)
			}

			if testList.Tail != lastNode {
				t.Fatal("The tail node must point to the last appended node")
			}
		},
	)

	t.Run(
		"Appending to a uninitialized list", func(t *testing.T) {
			testList := New()

			testList = nil

			_, err := testList.Append("blah", 12, "Lorem ipsum")

			if err == nil {
				t.Fatal("Expected action to return an error")
			}

			var want *UninitializedError

			if !errors.As(err, &want) {
				t.Fatalf("Expected error message: %q, got: %q", want, err)
			}
		},
	)
}

func TestList_Prepend(t *testing.T) {
	t.Run(
		"Prepend 100 times", func(t *testing.T) {
			testList := New()
			want := 100 // Expected list size

			for i := 0; i < want; i++ {

				_, err := testList.Prepend("Person", 12, "Passion")
				if err != nil {
					t.Fatal("An error occurred during testing")
				}
			}

			got := testList.Size

			if got != want {
				t.Fatalf("Expected size of the list to %d, but got %d", want, got)
			}
		},
	)

	t.Run(
		"Prepend to a list of two nodes", func(t *testing.T) {
			testList := New()

			testList.Prepend("Person 1", 12, "Passion 1")
			testList.Prepend("Person 2", 16, "Passion 2")
			newNode, _ := testList.Prepend("Person 3", 19, "Passion 3")

			want := "Person 3 (19, Passion 3)"
			got := newNode.Info()

			if got != want {
				t.Fatalf("Expected the node info to be %s, but got %s", want, got)
			}

			if testList.Head != newNode {
				t.Fatal("The head node must point to the last prepended node")
			}
		},
	)

	t.Run(
		"Prepending to a uninitialized list", func(t *testing.T) {
			testList := New()

			testList = nil

			_, err := testList.Prepend("blah", 12, "Lorem ipsum")

			if err == nil {
				t.Fatal("Expected action to return an error")
			}

			var want *UninitializedError

			if !errors.As(err, &want) {
				t.Fatalf("Expected error message: %q, got: %q", want, err)
			}
		},
	)
}

func TestList_AtIndex(t *testing.T) {
	t.Run(
		"Find index of an head node", func(t *testing.T) {
			testList := New()

			testList.Append("one", 1, "Passion 1")

			want := testList.Head
			got, err := testList.AtIndex(0)

			if err != nil {
				t.Fatal("Didn't expect an error to occur")
			}

			if got != want {
				t.Fatalf("Expected node data: %v, got: %v", want.Info(), got.Info())
			}
		},
	)

	t.Run(
		"Find index of an intermediate node", func(t *testing.T) {
			testList := New()

			testList.Append("lorem", 12, "ipsum")
			testList.Prepend("lorem2", 12, "ipsum2")

			intermediateNode, err := testList.InsertAt(1, "lorem3", 15, "ipsum3")

			if err != nil {
				t.Fatal("Insertion failed")
			}

			nodeFound, _ := testList.AtIndex(1)

			if nodeFound != intermediateNode {
				t.Fatalf("Expected node: %v, got: %v", intermediateNode, nodeFound)
			}
		},
	)

	t.Run(
		"Find index of the tail node", func(t *testing.T) {
			testList := New()

			testList.Append("one", 1, "Passion 1")
			tailNode, _ := testList.Append("one", 1, "Passion 1")

			want := tailNode
			got, err := testList.AtIndex(1)

			if err != nil {
				t.Fatal("Didn't expect an error to occur")
			}

			if got != want {
				t.Fatalf("Expected node data: %v, got: %v", want.Info(), got.Info())
			}
		},
	)

	t.Run(
		"invalid index", func(t *testing.T) {
			testList := New()

			testList.Append("lorem", 12, "ipsum")

			_, err := testList.AtIndex(10)

			if err == nil {
				t.Fatalf("Expected an error to be returned for invalid index")
			}
		},
	)
}

func TestList_DeleteAt(t *testing.T) {
	t.Run(
		"Delete the head node in a single-node list", func(t *testing.T) {
			testList := New()

			headNode, _ := testList.Append("lorem", 12, "ipsum")

			if testList.Size != 1 {
				t.Fatal("Expected insertion to have changed size from 0 to 1")
			}

			err := testList.DeleteAt(0)

			if err != nil {
				t.Fatalf("Expected deletion to work and not return error: %v", err)
			}

			if testList.Size != 0 {
				t.Fatal("Expected list size to change from 1 to 0 after deletion")
			}

			if testList.Head == headNode {
				t.Fatalf("Expected head node to be nil, not %v", headNode)
			}

			if testList.Tail != nil {
				t.Fatalf("Expexted tail node to be nil, nod %v", testList.Tail)
			}
		},
	)

	t.Run(
		"Deleting from an empty list", func(t *testing.T) {
			testList := New()

			for i := 0; i < 10; i++ {
				err := testList.DeleteAt(i)

				if err == nil {
					t.Fatal("Expected an error to be returned")
				}
			}
		},
	)

	t.Run(
		"Deleting with a non-existent index", func(t *testing.T) {
			testList := New()

			testList.Append("lorem", 12, "ipsum")

			err := testList.DeleteAt(4)

			if err == nil {
				t.Fatal("Expected an error to be returned")
			}
		},
	)

	t.Run(
		"A successful deletion of an intermediate node", func(t *testing.T) {
			testList := New()

			testList.Append("lorem1", 12, "ipsum1")
			testList.Append("lorem3", 12, "ipsum3")

			intermediateNode, _ := testList.InsertAt(1, "lorem2", 21, "ipsum2")

			initialListSize := testList.Size

			// before deletion, let's confirm the node is where it's supposed to be
			if testList.Head.Next != intermediateNode {
				t.Fatalf(
					"Expected the head node to point to %v not %v", intermediateNode,
					testList.Head,
				)
			}

			// now let's delete the node at index 1
			err := testList.DeleteAt(1)

			if err != nil {
				t.Fatal("Expected deletion to have worked and no error returned")
			}

			// now verify that the size of the list has been reduced
			if testList.Size != initialListSize-1 {
				t.Fatalf(
					"Expected list size to be %d not %d", testList.Size, initialListSize,
				)
			}

			// finally, verify that the head node points to the next node
			// of the deleted node
			want := intermediateNode.Next
			got := testList.Head.Next

			if got != want {
				t.Fatalf("Expected head node to be pointing to %q not %q", want, got)
			}
		},
	)
}

func TestList_InsertAt(t *testing.T) {
	t.Run(
		"Inserting at the head of the list", func(t *testing.T) {
			testList := New()

			testList.Append("lorem", 12, "ipsum")
			testList.Append("lorem2", 12, "ipsum2")

			newNode, err := testList.InsertAt(0, "lorem3", 15, "ipsum3")

			if err != nil {
				t.Fatal("Expected insertion to work")
			}

			if testList.Head != newNode {
				t.Fatalf("Expected head node to be %v not %v", newNode, testList.Head)
			}

			if testList.Head.Next != newNode.Next {
				t.Fatalf(
					"Expected head node to point to %v not %v", newNode.Next,
					testList.Head.Next,
				)
			}
		},
	)

	t.Run(
		"Inserting at the tail of the list", func(t *testing.T) {
			testList := New()

			testList.Append("lorem", 12, "ipsum")
			testList.Append("lorem2", 12, "ipsum2")

			tailNode, err := testList.InsertAt(2, "lorem3", 15, "ipsum3")

			if err != nil {
				t.Fatal("Expected insertion to work")
			}

			if testList.Tail != tailNode {
				t.Fatalf("Expected tail node to be %v not %v", tailNode, testList.Tail)
			}

			if testList.Tail.Next != nil {
				t.Fatalf(
					"Expected tail node to point to nil not %v", testList.Tail.Next,
				)
			}
		},
	)

	t.Run(
		"Inserting at an intermediate position", func(t *testing.T) {
			testList := New()

			testList.Append("lorem", 12, "ipsum")
			testList.Append("lorem2", 12, "ipsum2")

			intermediateNode, err := testList.InsertAt(1, "lorem3", 15, "ipsum3")

			if err != nil {
				t.Fatal("Expected insertion to work")
			}

			if testList.Head.Next != intermediateNode {
				t.Fatalf(
					"Expected head node to point to %v not %v", intermediateNode,
					testList.Head.Next,
				)
			}
		},
	)

	t.Run(
		"Inserting at a negative index", func(t *testing.T) {
			testList := New()

			_, err := testList.InsertAt(-4, "lorem3", 15, "ipsum3")

			if err == nil {
				t.Fatal("Expected insertion to fail")
			}
		},
	)

	t.Run(
		"Inserting at an index greater than size inserts at tail",
		func(t *testing.T) {
			testList := New()
			initialTail, _ := testList.Append("lorem1", 23, "ipsum1")

			newTail, err := testList.InsertAt(2, "lorem3", 15, "ipsum3")

			if err != nil {
				t.Fatal("Expected insertion to work")
			}

			if testList.Tail != newTail {
				t.Fatalf("Expected tail to be %v not %v", newTail, initialTail)
			}

			if initialTail.Next != newTail {
				t.Fatalf(
					"Expected the previous tail node to now be %v not %v", newTail,
					initialTail.Next,
				)
			}
		},
	)
}
