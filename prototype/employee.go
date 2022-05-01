package prototype

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address struct {
	Suite               int
	StreetAddress, City string
}

type Employee struct {
	Name   string
	Office Address
}

func (e *Employee) DeepCopy() *Employee {
	// note: no error handling below
	b := bytes.Buffer{}
	enc := gob.NewEncoder(&b)
	_ = enc.Encode(e)

	// peek into structure
	fmt.Println(b.String())

	dec := gob.NewDecoder(&b)
	result := Employee{}
	_ = dec.Decode(&result)
	return &result
}

// employee factory
// either a struct or some functions
var mainOfficeEmployee = Employee{
	"", Address{0, "123 East Dr", "London"}}

var auxOfficeEmployee = Employee{
	"", Address{0, "66 West Dr", "London"}}

// utility method for configuring Employee
func newEmployee(proto *Employee, name string, suite int) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite
	return result
}

func NewMainOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&mainOfficeEmployee, name, suite)
}

func NewAuxOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&auxOfficeEmployee, name, suite)
}
