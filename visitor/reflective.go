package visitor

import (
	"fmt"
	"strings"
)

type IExpression interface {
	Print(sb *strings.Builder)
}

type IDoubleExpression struct {
	value float64
}

func (d *IDoubleExpression) Print(sb *strings.Builder) {
	fmt.Fprintf(sb, "%g", d.value)
}

type IAddtionExpression struct {
	left, right IExpression
}

func (a *IAddtionExpression) Print(sb *strings.Builder) {
	sb.WriteString("(")
	a.left.Print(sb)
	sb.WriteString("+")
	a.right.Print(sb)
	sb.WriteString(")")
}
