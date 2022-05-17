package lsp

type AntiSquare struct {
	Rectangle
}

func NewAntiSquare(size int) *AntiSquare {
	return &AntiSquare{
		Rectangle: Rectangle{
			width:  size,
			height: size,
		},
	}
}

func (s *AntiSquare) GetWidth() int {
	return s.width
}

func (s *AntiSquare) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *AntiSquare) GetHeight() int {
	return s.height
}

func (s *AntiSquare) SetHeight(height int) {
	s.height = height
	s.width = height
}
