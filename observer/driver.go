package observer

import (
	"container/list"
)

type PropertyChanged struct {
	Name  string
	Value any
}

type Driver struct {
	Observable
	age int
}

func NewDriver(age int) *Driver {
	return &Driver{Observable{new(list.List)}, age}
}

func (d *Driver) Age() int {
	return d.age
}

func (d *Driver) SetAge(age int) {
	if age == d.age {
		return
	}
	d.age = age
	d.Fire(PropertyChanged{
		Name:  "Age",
		Value: d.age,
	})
}
