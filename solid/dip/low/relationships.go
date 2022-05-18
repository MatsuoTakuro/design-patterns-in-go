package low

import (
	"github.com/MatsuoTakuro/design-patterns-in-go/solid/dip/high"
)

type Relationships struct {
	relations []high.Info
}

func NewRelationships() *Relationships {
	return &Relationships{}
}

func (rs *Relationships) FindAllChildrenOf(parent *high.Person) []*high.Person {
	result := make([]*high.Person, 0)

	for i, v := range rs.relations {
		if v.Relationship() == high.Parent && v.From().Name() == parent.Name() {
			result = append(result, rs.relations[i].To())
		}
	}
	return result
}

func (rs *Relationships) AddParentAndChild(parent, child *high.Person) {
	rs.relations = append(rs.relations, *high.NewInfo(parent, high.Parent, child))
}
