package observer

import (
	"container/list"
)

type PropertyChanged struct {
	Name  string
	Value any
}

type Person struct {
	Observable
	age int
}

func NewPerson(age int) *Person {
	return &Person{Observable{new(list.List)}, age}
}

func (p *Person) Age() int {
	return p.age
}

func (p *Person) SetAge(age int) {
	if age == p.age {
		return
	}
	// canVote is depend on age.
	// instead of this property dependency, it may be better to build some sort of map
	// where all dependencies between all the different propertes are catalogued.
	oldCanVote := p.canVote()

	p.age = age
	p.Fire(PropertyChanged{
		Name:  "Age",
		Value: p.age,
	})

	if oldCanVote != p.canVote() {
		p.Fire(PropertyChanged{
			Name:  "CanVote",
			Value: p.canVote(),
		})
	}
}

func (p *Person) canVote() bool {
	return p.age >= 18
}
