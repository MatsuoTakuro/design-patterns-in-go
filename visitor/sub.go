package visitor

import (
	"fmt"
	"strings"
)

func Sub() {
	// intrusive()
	// reflective()
	classic()
}

func intrusive() {
	// 1+(2+3)
	e := IAddtionExpression{
		&IDoubleExpression{1}, // left
		&IAddtionExpression{
			left:  &IDoubleExpression{2}, // right >> left
			right: &IDoubleExpression{3}, // right >> right
		},
	}
	sb := strings.Builder{}
	e.Print(&sb)
	fmt.Println(sb.String())
}

func reflective() {
	// 1+(2+3)
	e := &RAddtionExpression{
		&RDoubleExpression{1}, // left
		&RAddtionExpression{
			left:  &RDoubleExpression{2}, // right >> left
			right: &RDoubleExpression{3}, // right >> right
		},
	}
	sb := strings.Builder{}
	Print(e, &sb)
	fmt.Println(sb.String())
}

func classic() {
	// 1+(2+3)
	d1 := &DoubleExpression{1}
	d2 := &DoubleExpression{2}
	d3 := &DoubleExpression{3}
	a := &AdditionExpression{
		left:  d2,
		right: d3,
	}
	e := &AdditionExpression{
		left:  d1,
		right: a,
	}

	ep := NewExpressionPrinter() // visitor
	ep.VisitDoubleExpression(d1) // ep is already accepted by d1
	fmt.Println(ep)

	ep1 := NewExpressionPrinter()  // visitor
	ep1.VisitAdditionExpression(a) // ep1 is already accepted by a
	fmt.Println(ep1)

	ep2 := NewExpressionPrinter()  // visitor
	ep2.VisitAdditionExpression(e) // ep2 is already accepted by e
	fmt.Println(ep2)
}
