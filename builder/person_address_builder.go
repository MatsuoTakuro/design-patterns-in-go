package builder

type PersonAddressBuilder struct {
	PersonBuilder
}

func (pab *PersonAddressBuilder) At(streetAdress string) *PersonAddressBuilder {
	pab.person.streetAddress = streetAdress
	return pab
}

func (pab *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	pab.person.city = city
	return pab
}

func (pab *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	pab.person.postcode = postcode
	return pab
}
