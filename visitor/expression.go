package visitor

type Expression interface {
	Accept(ev ExpressionVisitor) // accept a visitor
}
