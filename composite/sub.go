package composite

import "fmt"

func Sub() {
	drawring := GraphicObject{"My Drawing", "", nil}
	drawring.Children = append(drawring.Children, *NewSquare("Red"))
	drawring.Children = append(drawring.Children, *NewCircle("Yellow"))

	group := GraphicObject{"Group 1", "", nil}
	group.Children = append(group.Children, *NewCircle("Blue"))
	group.Children = append(group.Children, *NewSquare("Blue"))
	drawring.Children = append(drawring.Children, group)

	fmt.Println(drawring.String())
}
