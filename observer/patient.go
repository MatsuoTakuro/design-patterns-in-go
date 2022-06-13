package observer

import "container/list"

// whenever a person catches a cold,
// a doctor must be called
type Patient struct {
	Observable
	Name string
}

func NewPatient(name string) *Patient {
	return &Patient{
		Observable: Observable{
			subs: new(list.List),
		},
		Name: name,
	}
}

func (p *Patient) CatchACold() {
	p.Fire(p.Name)
}
