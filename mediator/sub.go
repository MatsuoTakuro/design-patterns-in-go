package mediator

import "fmt"

func Sub() {
	c := Channel{
		people: []*Member{},
	}
	john := NewMember("John")
	jane := NewMember("Jane")

	c.Join(john)
	c.Join(jane)

	john.SayHere("Hi, channel")
	jane.SayHere("Oh, hey John")

	simon := NewMember("Simon")
	c.Join(simon)
	simon.SayHere("Hi, everyone!")

	jane.DirectMessage("Simon", "Glad you could join us!")

	fmt.Println(c)
	fmt.Println(john)
	fmt.Println(jane)
	fmt.Println(simon)
}
