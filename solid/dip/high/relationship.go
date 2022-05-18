package high

type Relationship int

const (
	Parent Relationship = iota + 1
	Child
	Sibling
)
