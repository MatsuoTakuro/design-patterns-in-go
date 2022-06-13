package observer

import "fmt"

func Sub() {
	// doctorService()
	trafficManagement()
}

func doctorService() {
	p := NewPatient("Boris")
	ds := &DoctorService{}
	p.Subscribe(ds)

	// let's test it!
	p.CatchACold()
}

func trafficManagement() {
	d := NewDriver(15)
	t := &TrafficManagement{
		Observable: d.Observable,
	}
	d.Subscribe(t)

	for i := 16; i <= 20; i++ {
		fmt.Println("Setting age to ", i)
		d.SetAge(i)
	}

	fmt.Println(d.Age())
}
