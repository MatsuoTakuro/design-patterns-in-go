package lsp

type Square struct {
	size int
}

func (s Square) convRectangle() *Rectangle {
	return &Rectangle{
		width:  s.size,
		height: s.size,
	}
}
