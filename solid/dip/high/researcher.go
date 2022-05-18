package high

import (
	"fmt"
)

type Researcher struct {
	browser RelationshipBrowser
}

func NewResearcher(browser RelationshipBrowser) *Researcher {
	return &Researcher{
		browser: browser,
	}
}

func (r *Researcher) Investigate(target Person) {
	for _, p := range r.browser.FindAllChildrenOf(&target) {
		fmt.Printf("%s has a child called %s.\n", target.Name(), p.Name())
	}
}
