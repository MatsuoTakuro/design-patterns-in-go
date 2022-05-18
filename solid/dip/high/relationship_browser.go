package high

type RelationshipBrowser interface {
	FindAllChildrenOf(parent *Person) []*Person
}
