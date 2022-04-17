package iterator

type PersonNameIterator struct {
	person       *Person
	currentIndex int
}

func NewPersonNameIterator(person *Person) *PersonNameIterator {
	return &PersonNameIterator{person, -1}
}

func (it *PersonNameIterator) HasNext() bool {
	it.currentIndex++
	return it.currentIndex < 3
}

func (it *PersonNameIterator) Name() string {
	switch it.currentIndex {
	case 0:
		return it.person.Firstname
	case 1:
		return it.person.Middlename
	case 2:
		return it.person.Lastname
	}
	panic("We should not be here!")
}
