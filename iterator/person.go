package iterator

type Person struct {
	Firstname, Middlename, Lastname string
}

func (p *Person) Names() [3]string {
	return [3]string{p.Firstname, p.Middlename, p.Lastname}
}

func (p *Person) NamesGenerator() <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		out <- p.Firstname
		if len(p.Middlename) > 0 {
			out <- p.Middlename
		}
		out <- p.Lastname
	}()
	return out
}
