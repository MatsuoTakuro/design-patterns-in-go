package composite

import "fmt"

func Sub() {
	// graphicObject()
	neuralNetworks()
}

func graphicObject() {
	drawring := GraphicObject{"My Drawing", "", nil}
	drawring.Children = append(drawring.Children, *NewSquare("Red"))
	drawring.Children = append(drawring.Children, *NewCircle("Yellow"))

	group := GraphicObject{"Group 1", "", nil}
	group.Children = append(group.Children, *NewCircle("Blue"))
	group.Children = append(group.Children, *NewSquare("Blue"))
	drawring.Children = append(drawring.Children, group)

	fmt.Println(drawring.String())
}

func neuralNetworks() {
	neuron1, neuron2 := NewNeuron(), NewNeuron()
	layer1, layer2 := NewNeuronLayer(3), NewNeuronLayer(4)

	Connect(neuron1, neuron2)
	Connect(neuron1, layer1)
	Connect(layer2, neuron1)
	Connect(layer1, layer2)

	neuron1.printName()
	neuron2.printName()
	layer1.printNames()
	layer2.printNames()
}
