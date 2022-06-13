package observer

func Sub() {
	p := NewPatient("Boris")
	ds := &DoctorService{}
	p.Subscribe(ds)

	// let's test it!
	p.CatchACold()
}
