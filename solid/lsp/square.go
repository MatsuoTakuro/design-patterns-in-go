package lsp

// modified LSP
// If a function takes an interface(Sized) and
// works with a type T(Square) that implements this interface(Sized),
// any structure(for example, Rectangle) that aggregates T(Square)
// should also be usable in that function.

type Square struct {
	size int
}

// No need to forcibly implement the Sized interface in the same way as Rectangle.
// Just, Rectangle can aggregate Square.
func (s Square) convRectangle() *Rectangle {
	return &Rectangle{
		width:  s.size,
		height: s.size,
	}
}
