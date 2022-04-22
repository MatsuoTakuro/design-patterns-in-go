package factories

import "fmt"

type tiredPerson struct {
	name string
	age  int
}

func (tp *tiredPerson) SayHello() {
	fmt.Println("Sorry, I am too tired.")
}
