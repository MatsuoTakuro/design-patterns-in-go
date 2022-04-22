package factories

import "fmt"

func Sub() {
	factoryFunctionAndInterfaceFactory()
	factoryGeneratorWithFunction()
	factoryGeneratorWithStruct()
	prototypeFactory()
}

func factoryFunctionAndInterfaceFactory() {
	p := NewPerson("John", 33)
	p.SayHello()
	tp := NewPerson("James", 133)
	tp.SayHello()

}

func factoryGeneratorWithFunction() {
	developerFactory := NewEmployeeFactory("Developer", 60000)
	managerFactory := NewEmployeeFactory("Manager", 80000)

	developer := developerFactory("Adam")
	manager := managerFactory("Jane")

	fmt.Println(*developer)
	fmt.Println(*manager)
}

func factoryGeneratorWithStruct() {
	bossFactory := NewEmployeeFactory2("CEO", 100000)
	bossFactory.AnnualIncome = 11000
	boss := bossFactory.Create("Sam")
	fmt.Println(*boss)
}

func prototypeFactory() {
	m := NewEmployee(Manager)
	m.Name = "Sam"
	fmt.Println(*m)
}
