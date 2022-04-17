package iterator

type InOrderIterator struct {
	Current         *Node
	root            *Node
	traversalStared bool
}

func NewInOrderIterator(root *Node) *InOrderIterator {
	it := &InOrderIterator{
		Current:         root,
		root:            root,
		traversalStared: false}
	for it.Current.left != nil {
		it.Current = it.Current.left
	}
	return it
}

func (it *InOrderIterator) Reset() {
	it.Current = it.root
	it.traversalStared = false
}

func (it *InOrderIterator) HasNext() bool {

	if it.Current == nil {
		return false
	}

	if !it.traversalStared {
		it.traversalStared = true
		return it.traversalStared // can use first element
	}

	if it.Current.right != nil {
		it.Current = it.Current.right
		for it.Current.left != nil {
			it.Current = it.Current.left
		}
		return true

	} else {

		p := it.Current.parent
		for p != nil && it.Current == p.right {
			it.Current = p
			p = p.parent
		}
		it.Current = p
		return it.Current != nil
	}
}
