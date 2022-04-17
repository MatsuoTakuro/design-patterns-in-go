package iterator

import "fmt"

func Sub2() {
	//    1
	//   / \
	//  2   3

	// binary-tree-traversal
	// in-order:  213
	// preorder:  123
	// postorder: 231

	root := NewNode(1, NewTerminalNode(2), NewTerminalNode(3))
	it := NewInOrderIterator(root)
	for it.HasNext() {
		fmt.Printf("%d,", it.Current.Number)
	}
	fmt.Println("\b")

	t := NewBinaryTree(root)
	for it := t.InOrder(); it.HasNext(); {
		fmt.Printf("%d,", it.Current.Number)
	}
	fmt.Println("\b")
}
