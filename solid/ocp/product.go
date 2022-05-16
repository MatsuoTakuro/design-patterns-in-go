package ocp

type Color int

const (
	red Color = iota + 1
	green
	blue
)

type Size int

const (
	small Size = iota + 1
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}
