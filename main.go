package main

import (
	"fmt"

	singleResponsibility "github.com/MatsuoTakuro/design-patterns-in-go/solid_design_principles/single_responsibility"
)

func main() {
	j := singleResponsibility.Journal{}
	j.AddEntry("I cried today")
	j.AddEntry("I ate a bug")
	fmt.Println(j.String())

	// anti pattern
	singleResponsibility.SaveToFile(&j, "solid_design_principles/single_responsibility/journal.txt")

	p := singleResponsibility.Persistence{LineSeparator: "\r\n"}
	p.SaveToFile(&j, "solid_design_principles/single_responsibility/journal2.txt")

}
