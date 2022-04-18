package strategy

import "strings"

type HtmlListStrategy struct{}

func (m *HtmlListStrategy) Start(builder *strings.Builder) {
	builder.WriteString("<ul>\n")
}

func (m *HtmlListStrategy) End(builder *strings.Builder) {
	builder.WriteString("</ul>\n")
}

func (m *HtmlListStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString(" <li> " + item + "</li>\n")
}
