package node

import (
	"testing"
)

type nodeTest struct {
	name    string
	age     int
	passion string
}

var nodeTests = []nodeTest{
	{"John Doe", 45, "To be successful"},
	{"Sally Banks", 12, "To be a musician"},
}

func TestCreate(t *testing.T) {
	t.Run(
		"Creating new nodes", func(t *testing.T) {
			for _, test := range nodeTests {
				newNode := Create(test.name, test.age, test.passion)

				if newNode == nil {
					t.Fatalf("The %q function failed to create a node", "Create")
				}

				if newNode.Name != test.name {
					t.Fatalf("expected %v, got %v", test.name, newNode.Name)
				}

				if newNode.Age != test.age {
					t.Fatalf("expected %v, got %v", test.age, newNode.Age)
				}

				if newNode.Next != nil {
					t.Fatal("A new node cannot have a next node by default")
				}
			}
		},
	)

	t.Run(
		"Adding a next node to an existing node", func(t *testing.T) {
			newNode := Create("John Doe", 45, "To be a good man")
			anotherNode := Create("Sally Banks", 45, "To live")

			newNode.Next = anotherNode

			if newNode.Next.Name != anotherNode.Name {
				t.Fatalf("Expected %v, got %v", anotherNode.Name, newNode.Next.Name)
			}
		},
	)
}

func TestInfo(t *testing.T) {
	t.Run(
		"Test info after creating a node", func(t *testing.T) {
			newNode := Create("John Doe", 45, "To be a good man")

			want := "John Doe (45, To be a good man)"
			got := newNode.Info()

			if got != want {
				t.Fatalf("Expected %v, got %v", want, got)
			}
		},
	)

	t.Run(
		"Update and verify node info", func(t *testing.T) {
			newNode := Create("John Doe", 45, "To be a good man")

			if newNode.Name != "John Doe" {
				t.Fatalf("The name of the user must be %v after creation", "John Doe")
			}

			newName := "Sally Banks"
			newNode.Name = newName

			if newNode.Name != newName {
				t.Fatalf(
					"Expected name to have changed from %q to %q", newNode.Name, newName,
				)
			}
		},
	)
}

func BenchmarkCreate(b *testing.B) {
	b.Run(
		"When only one is created multiple times", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				Create("John Doe", 45, "To be a good man")
			}
		},
	)

	b.Run(
		"When more than one nodes is created at about the same time concurrently",
		func(b *testing.B) {

			go func() {
				for n := 0; n < b.N; n++ {
					Create("John Doe", 45, "To be a good man")
					Create("Sally", 25, "To be a good woman")
					Create("Bob", 12, "To be a good child")
				}
			}()
		},
	)

	b.Run(
		"When more than one nodes is created at about the same time without concurrency",
		func(b *testing.B) {

			go func() {
				for n := 0; n < b.N; n++ {
					Create("John Doe", 45, "To be a good man")
					Create("Sally", 25, "To be a good woman")
					Create("Bob", 12, "To be a good child")
				}
			}()
		},
	)
}
