package lsp

import "fmt"

func Sub() {
	rc := &Rectangle{
		width:  2,
		height: 3,
	}
	UseIt(rc)

	antiSq := NewAntiSquare(5)
	UseIt(antiSq)

	sq := Square{
		size: 5,
	}
	// Square follows LSP!
	rc2 := sq.convRectangle()
	UseIt(rc2)
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Print("Expected an area of ", expectedArea, ", but got ", actualArea, " width: ", sized.GetWidth(), " heigt: ", sized.GetHeight(), "\n")
}
