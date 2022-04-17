package iterator

type Node struct {
	Number              int
	left, right, parent *Node
}

func NewNode(number int, left *Node, right *Node) *Node {
	n := &Node{
		Number: number,
		left:   left,
		right:  right}
	left.parent = n
	right.parent = n
	return n
}

func NewTerminalNode(number int) *Node {
	return &Node{Number: number}
}
