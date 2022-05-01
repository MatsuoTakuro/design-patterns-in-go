package prototype

import "fmt"

func Sub() {

	// john := &Person{
	// 	"John",
	// 	&PersonAddress{"123 London Rd", "London", "UK"},
	// 	[]string{"Chris", "Matt"}}
	// jane := john.DeepCopy()
	// jane.Name = "Jane"
	// jane.Address.StreetAddress = "321 Baker St"
	// jane.Friends = append(jane.Friends, "Angela")

	// fmt.Println(*john, *john.Address)
	// fmt.Println(*jane, *jane.Address)
	john := NewMainOfficeEmployee("John", 100)
	jane := NewAuxOfficeEmployee("Jane", 200)
	fmt.Println(*john)
	fmt.Println(*jane)

}
