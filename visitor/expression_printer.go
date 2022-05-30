package visitor

import (
	"fmt"
	"strings"
)

type ExpressionVisitor interface {
	VisitDoubleExpression(de *DoubleExpression)
	VisitAdditionExpression(ae *AdditionExpression)
}

type ExpressionPrinter struct { // visitor
	sb strings.Builder
}

// visitor executes this if accepted.
func (e *ExpressionPrinter) VisitDoubleExpression(de *DoubleExpression) {
	e.sb.WriteString(fmt.Sprintf("%g", de.value))
}

// visitor executes this if accepted.
func (e *ExpressionPrinter) VisitAdditionExpression(ae *AdditionExpression) {
	e.sb.WriteString("(")
	ae.left.Accept(e)
	e.sb.WriteString("+")
	ae.right.Accept(e)
	e.sb.WriteString(")")
}

func (e ExpressionPrinter) String() string {
	return e.sb.String()
}

func NewExpressionPrinter() *ExpressionPrinter {
	return &ExpressionPrinter{
		sb: strings.Builder{},
	}
}
