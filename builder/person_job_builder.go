package builder

type PersonJobBuilder struct {
	PersonBuilder
}

func (pjb *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	pjb.person.companyName = companyName
	return pjb
}

func (pjb *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	pjb.person.position = position
	return pjb
}

func (pjb *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	pjb.person.annualIncome = annualIncome
	return pjb
}
