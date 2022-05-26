package proxy

func Sub() {
	protectionProxy()
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
