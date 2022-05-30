package visitor

import (
	"fmt"
	"strings"
)

type RExpression interface {
	// nothing here!
}

type RDoubleExpression struct {
	value float64
}

type RAddtionExpression struct {
	left, right RExpression
}

func Print(e RExpression, sb *strings.Builder) {
	if de, ok := e.(*RDoubleExpression); ok {
		fmt.Fprintf(sb, "%g", de.value)
	} else if ae, ok := e.(*RAddtionExpression); ok {
		sb.WriteString("(")
		Print(ae.left, sb)
		sb.WriteString("+")
		Print(ae.right, sb)
		sb.WriteString(")")
	}
	// breaks OCP
	// will work incorrectly on missing case
}
