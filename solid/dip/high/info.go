package high

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

func (i Info) From() *Person {
	return i.from
}

func (i Info) Relationship() Relationship {
	return i.relationship
}

func (i Info) To() *Person {
	return i.to
}

func NewInfo(from *Person, relationship Relationship, to *Person) *Info {
	return &Info{
		from: &Person{
			name: from.Name(),
		},
		relationship: relationship,
		to: &Person{
			name: to.Name(),
		},
	}
}
