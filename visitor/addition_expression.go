package visitor

type AdditionExpression struct {
	left, right Expression
}

func (ae *AdditionExpression) Accept(ev ExpressionVisitor) {
	ev.VisitAdditionExpression(ae)
}
