package ocp

import (
	"fmt"
)

func Sub() {
	apple := Product{
		name:  "Apple",
		color: green,
		size:  small,
	}
	tree := Product{
		name:  "Tree",
		color: green,
		size:  large,
	}
	house := Product{
		name:  "House",
		color: blue,
		size:  large,
	}
	products := []Product{apple, tree, house}
	f := Filter{}

	fmt.Println("Green products:")
	greenSpec := ColorSpecification{
		color: green,
	}
	for _, v := range f.by(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	largeSpec := SizeSpecification{
		size: large,
	}
	largeGreenSpec := AndSpecification{
		first:  largeSpec,
		second: greenSpec,
	}
	fmt.Println("Large green items:")
	for _, v := range f.by(products, largeGreenSpec) {
		fmt.Printf(" - %s is large and green\n", v.name)
	}
}
