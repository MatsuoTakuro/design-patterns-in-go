package proxy

import "fmt"

type CarProxy struct {
	car    Car
	driver *Driver
}

func (cp *CarProxy) Drive() {
	if cp.driver.Age >= 16 {
		cp.car.Drive()
	} else {
		fmt.Println("Driver too young.")
	}
}

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{
		car:    Car{},
		driver: driver}
}
