package factories

import "fmt"

func Sub() {
	p := NewPerson("John", 33)
	fmt.Println(*p)
}
