package builder

type PersonBuilder struct {
	person *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

func (pb *PersonBuilder) Build() *Person {
	return pb.person
}

func (pb *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*pb}
}

func (pb *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*pb}
}
