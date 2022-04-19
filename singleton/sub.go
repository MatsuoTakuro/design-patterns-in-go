package singleton

import "fmt"

func Sub() {
	db := GetSingletonDatabase()
	pop := db.GetPopulation("Seoul")
	fmt.Println("Pop of Seoul is ", pop)
}
