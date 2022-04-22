package factories

import "fmt"

func Sub() {
	factoryFunctionAndInterfaceFactory()
	factoryGenerator()
}

func factoryFunctionAndInterfaceFactory() {
	p := NewPerson("John", 33)
	p.SayHello()
	tp := NewPerson("James", 133)
	tp.SayHello()

}

func factoryGenerator() {
	developerFactory := NewEmployeeFactory("developer", 60000)
	managerFactory := NewEmployeeFactory("manager", 80000)

	developer := developerFactory("Adam")
	manager := managerFactory("Jane")

	fmt.Println(*developer)
	fmt.Println(*manager)
}
