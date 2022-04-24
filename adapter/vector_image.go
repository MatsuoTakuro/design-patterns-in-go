package adapter

type Line struct {
	X1, Y1, X2, Y2 int
}

type VectorImage struct {
	Lines []Line
}

func NewRectangle(width, height int) *VectorImage {
	width -= 1
	height -= 1
	return &VectorImage{[]Line{
		{0, 0, width, 0},           // bottom
		{0, 0, 0, height},          // left
		{width, 0, width, height},  // right
		{0, height, width, height}, // top
	}}
}
