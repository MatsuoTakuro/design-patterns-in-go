package ocp

type Filter struct{}

func (f Filter) by(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	// Bad Loop
	// for _, v := range products {
	// 	if spec.IsSatisfied(&v) {
	// 		result = append(result, &v) // For pointer values, all assignments to v are overwritten because of reference passing.
	// 	}
	// }
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}
