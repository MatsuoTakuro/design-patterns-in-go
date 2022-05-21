package decorator

import "fmt"

type Lizard struct {
	age int
}

func (l *Lizard) Age() int {
	return l.age
}

func (l *Lizard) SetAge(age int) {
	l.age = age
}

func (l *Lizard) Crawl() {
	if l.age < 10 {
		fmt.Println("Crawling!")
	}
}
