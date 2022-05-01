package prototype

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type PersonAddress struct {
	StreetAddress, City, Country string
}
type Person struct {
	Name    string
	Address *PersonAddress
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	// note: no error handling below
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	// peek into structure
	fmt.Println(b.String())

	d := gob.NewDecoder(&b)
	result := Person{}
	_ = d.Decode(&result)
	return &result
}
