package adapter

import "fmt"

func Sub() {
	rc := NewRectangle(6, 4)
	a := VectorToRaster(rc)
	fmt.Print(DrawPoints(a))
}
