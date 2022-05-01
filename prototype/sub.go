package prototype

import "fmt"

func Sub() {

	john := Person{"John", &Address{"123 London Rd", "London", "UK"}}
	jane := john
	jane.Name = "Jane"
	jane.Address = &Address{
		john.Address.StreetAddress,
		john.Address.City,
		john.Address.Country}
	// jane.Address.StreetAddress = "321 Baker St"
	jane.Address.StreetAddress = "321 Baker St"

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)

}
