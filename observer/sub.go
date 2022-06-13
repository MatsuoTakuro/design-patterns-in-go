package observer

import "fmt"

func Sub() {
	// doctorService()
	// trafficManagement()
	canVote()
}

func doctorService() {
	p := NewPatient("Boris")
	ds := &DoctorService{}
	p.Subscribe(ds)

	// let's test it!
	p.CatchACold()
}

func trafficManagement() {
	d := NewPerson(15)
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

func canVote() {
	p := NewPerson(0)
	er := &ElectrocalRoll{
		Observable: p.Observable,
	}
	p.Subscribe(er)

	for i := 10; i < 20; i++ {
		fmt.Println("Setting age to", i)
		p.SetAge(i)
	}
}
