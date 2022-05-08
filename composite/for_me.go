package composite

import "fmt"

var nNum int = 1
var lNum int = 1

func (n *Neuron) printName() {
	fmt.Printf("Neuron: %s\n", n.name)
	fmt.Print("In    : ")
	for _, n := range n.In {
		fmt.Print(n.name, ", ")
	}
	fmt.Println()
	fmt.Print("Out   : ")
	for _, n := range n.Out {
		fmt.Print(n.name, ", ")
	}
	fmt.Println()
	fmt.Println()
}

func (nl *NeuronLayer) printNames() {
	fmt.Printf("NeuronLayer: %s\n", nl.name)
	for _, n := range nl.Neurons {
		n.printName()
	}
	fmt.Println()
}
