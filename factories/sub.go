package factories

func Sub() {
	p := NewPerson("John", 33)
	p.SayHello()
	tp := NewPerson("James", 133)
	tp.SayHello()
}
