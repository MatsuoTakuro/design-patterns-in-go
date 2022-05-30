package visitor

type DoubleExpression struct {
	value float64
}

func (de *DoubleExpression) Accept(ev ExpressionVisitor) {
	ev.VisitDoubleExpression(de)
}
