package proxy

func Sub() {
	// protectionProxy()
	virtualProxy()
}

func protectionProxy() {
	car := NewCarProxy(&Driver{
		Age: 12,
	})
	car.Drive()
	car2 := NewCarProxy(&Driver{
		Age: 17,
	})
	car2.Drive()
}

func virtualProxy() {
	bmp := NewLazyBitmap("demo.png")
	DrawImage(bmp)
	DrawImage(bmp)
}
