package strategy

import "fmt"

func Sub() {
	tp := NewTextProcessor(&MarkdownListStrategy{})
	tp.AppendList([]string{"foo", "bar", "baz"})
	fmt.Println(tp)

	tp = NewTextProcessor(&HtmlListStrategy{})
	tp.AppendList([]string{"foo", "bar", "baz"})
	fmt.Println(tp)
}
