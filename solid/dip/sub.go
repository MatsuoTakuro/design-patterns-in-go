package dip

import (
	"github.com/MatsuoTakuro/design-patterns-in-go/solid/dip/high"
	"github.com/MatsuoTakuro/design-patterns-in-go/solid/dip/low"
)

//* Dependency Inversion Principle
//* High-level module(HLM) should not depend on low-level module(LLM)
//* Both should depend on abstractions
// TODO: But, can a low-level module depend on a high-level module, not only 'abstractions'?

func Sub() {
	//* high-level module
	parent := high.NewPerson("John")
	child1 := high.NewPerson("Chris")
	child2 := high.NewPerson("Matt")

	//* low-level module
	rs := low.NewRelationships()
	rs.AddParentAndChild(parent, child1)
	rs.AddParentAndChild(parent, child2)

	//* high-level module
	rch := high.NewResearcher(rs)
	rch.Investigate(*parent)
}
