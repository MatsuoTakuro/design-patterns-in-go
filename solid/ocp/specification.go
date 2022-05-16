package ocp

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return c.color == p.color
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return s.size == p.size
}

type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}
