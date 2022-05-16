package bridge

func Sub() {
	rr := RasterRenderer{
		Dpi: 0,
	}
	c := NewCircle(&rr, 5)
	c.Draw()
	c.Resize(2)
	c.Draw()

	vr := VectorRenderer{}
	c2 := NewCircle(&vr, 5)
	c2.Draw()
	c2.Resize(2)
	c2.Draw()
}
