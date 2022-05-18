package isp

func Sub() {
	d := &Document{
		name: "GoF Design Patterns",
	}

	myp := &MyPrinter{}
	myp.Print(*d)

	pc := &Photocopier{}
	pc.Print(*d)
	pc.Scan(*d)

	Print(pc, *d)
	Scan(pc, *d)

	mfm := &MultiFunctionMachine{
		printer: myp,
		scanner: pc,
	}
	mfm.Print(*d)
	mfm.Scan(*d)
}
