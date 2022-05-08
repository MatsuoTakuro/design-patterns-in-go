package composite

import "strconv"

type Neuron struct {
	name    string
	In, Out []*Neuron
}

func NewNeuron() *Neuron {
	name := "n" + strconv.Itoa(nNum)
	nNum++
	return &Neuron{name: name}
}

func (n *Neuron) Iter() []*Neuron {
	return []*Neuron{n}
}

func (n *Neuron) ConnectTo(other *Neuron) {
	n.Out = append(n.Out, other)
	other.In = append(other.In, n)
}
