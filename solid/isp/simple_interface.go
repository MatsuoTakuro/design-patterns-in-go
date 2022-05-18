package isp

type Document struct {
	name string
}

func (d Document) String() string {
	return d.name
}

// * Better approach: split into several interfaces *
type Printer interface {
	Print(d Document)
}

type Fax interface {
	Fax(d Document)
}

type Scanner interface {
	Scan(d Document)
}
