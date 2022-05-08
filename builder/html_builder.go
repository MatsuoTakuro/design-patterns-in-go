package builder

type HtmlBuilder struct {
	rootName string
	root     HtmlElement
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	b := HtmlBuilder{rootName, HtmlElement{rootName, "", []HtmlElement{}}}
	return &b
}

func (b *HtmlBuilder) String() string {
	return b.root.String()
}

func (b *HtmlBuilder) AddChild(ChildName, childText string) {
	e := HtmlElement{ChildName, childText, []HtmlElement{}}
	b.root.elements = append(b.root.elements, e)
}

func (b *HtmlBuilder) AddChildFluent(ChildName, childText string) *HtmlBuilder {
	e := HtmlElement{ChildName, childText, []HtmlElement{}}
	b.root.elements = append(b.root.elements, e)
	return b
}
