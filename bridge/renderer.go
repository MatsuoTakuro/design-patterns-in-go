package bridge

import "fmt"

type Renderer interface {
	RenderCircle(radius float32)
}

type VectorRenderer struct{}

func (v VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing lines for a circle of radius", radius)
}

type RasterRenderer struct {
	Dpi int
}

func (r RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing pixels for a circle of radius", radius)
}
