package visitor

type Expression interface {
	Accept(ev ExpressionVisitor)
}
