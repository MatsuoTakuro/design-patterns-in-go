package srp

import "fmt"

func Sub() {
	j := Journal{}
	j.AddEntry("I cried today")
	j.AddEntry("I ate a bug")
	fmt.Println(j.String())

	// anti pattern
	SaveToFile(&j, "solid_design_principles/srp/journal.txt")

	p := Persistence{LineSeparator: "\r\n"}
	p.SaveToFile(&j, "solid_design_principles/srp/journal2.txt")
}
