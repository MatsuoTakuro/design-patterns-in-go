package decorator

func Sub() {
	dragon()

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
