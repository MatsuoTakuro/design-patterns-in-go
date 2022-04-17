package iterator

import "fmt"

func Sub1() {
	p := Person{
		Firstname:  "Alexander",
		Middlename: "Graham",
		Lastname:   "Bell"}

	for _, name := range p.Names() {
		fmt.Println(name)
	}

	for name := range p.NamesGenerator() {
		fmt.Println(name)
	}

	for it := NewPersonNameIterator(&p); it.HasNext(); {
		fmt.Println(it.Name())
	}
}
