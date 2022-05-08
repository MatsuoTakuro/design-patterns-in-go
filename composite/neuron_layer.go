package composite

import "strconv"

type NeuronLayer struct {
	name    string
	Neurons []Neuron
}

func (nl *NeuronLayer) Iter() []*Neuron {
	result := make([]*Neuron, 0)
	for i := range nl.Neurons {
		result = append(result, &nl.Neurons[i])
	}
	return result
}

// func NewNeuronLayer(count int) *NeuronLayer {
// 	return &NeuronLayer{make([]Neuron, count)}
// }

func NewNeuronLayer(count int) *NeuronLayer {
	var neurons []Neuron
	name := "l" + strconv.Itoa(lNum)
	lNum++
	for _, n := range make([]Neuron, count) {
		n.name = "n" + strconv.Itoa(nNum)
		neurons = append(neurons, n)
		nNum++
	}
	return &NeuronLayer{name, neurons}
}
