package builder

import (
	"fmt"
	"strings"
)

func Sub() {
	helloWorld()
	helloWorldByBuilder()
	helloWorldByBuilderWithFluent()
}

func helloWorld() {
	hello := "hello"
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(hello)
	sb.WriteString("</p>")
	fmt.Printf("%s\n", sb.String())

	words := []string{"hello", "world"}
	sb.Reset()
	sb.WriteString("<ul>")
	for _, v := range words {
		sb.WriteString("<li>")
		sb.WriteString(v)
		sb.WriteString("</li>")
	}
	sb.WriteString("</ul>")
	fmt.Printf("%s\n", sb.String())

}

func helloWorldByBuilder() {
	b := NewHtmlBuilder("ul")
	b.AddChild("li", "hello")
	b.AddChild("li", "world")
	fmt.Println(b.String())
}

func helloWorldByBuilderWithFluent() {
	b := NewHtmlBuilder("ul")
	b.AddChildFluent("li", "hello").
		AddChildFluent("li", "world")
	fmt.Println(b.String())
}
