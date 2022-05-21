package decorator

type Dragon struct {
	bird   Bird
	lizard Lizard
}

func (d *Dragon) Age() int {
	return d.bird.age
}

func (d *Dragon) SetAge(age int) {
	d.bird.SetAge(age)
	d.lizard.SetAge(age)
}

func (d *Dragon) Fly() {
	d.bird.Fly()
}

func (d *Dragon) Crawl() {
	d.lizard.Crawl()
}

func NewDragon() *Dragon {
	return &Dragon{
		bird: Bird{
			age: 0,
		},
		lizard: Lizard{
			age: 0,
		},
	}
}
