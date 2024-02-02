package main

func main() {

}

type Node[T comparable] struct {
	Value  T
	Left   *Node
	Right  *Node
	Parent *Node
}

func NewGraph() *Node {
	return &Node{
		Left:   nil,
		Right:  nil,
		Parent: nil,
	}
}
