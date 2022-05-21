package decorator

import "fmt"

func Sub() {
	// dragon()
	shape()
}

func dragon() {
	d := NewDragon()

	d.SetAge(5)
	d.Fly()
	d.Crawl()

	d.SetAge(10)
	d.Fly()
	d.Crawl()
}

func shape() {
	c := Circle{
		Radius: 2.0,
	}
	fmt.Println(c.Render())

	redC := ColoredShape{
		Shape: &c,
		Color: "Red",
	}
	fmt.Println(redC.Render())

	rhsC := TransparentShape{
		Shape:        &redC,
		Transparency: 0.5,
	}
	fmt.Println(rhsC.Render())

	var x float32 = 2
	c.Resize(x)
	fmt.Println(c.Render())
	fmt.Println(redC.Render()) // affected by c.Resize(x)
	fmt.Println(rhsC.Render()) // affected by c.Resize(x)
}
