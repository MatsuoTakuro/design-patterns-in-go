package proxy

import "fmt"

type Car struct{}

func (c *Car) Drive() {
	fmt.Println("Car being driven.")
}
