package high

type Person struct {
	name string
}

func NewPerson(name string) *Person {
	return &Person{
		name: name,
	}
}

func (p Person) Name() string {
	return p.name
}
