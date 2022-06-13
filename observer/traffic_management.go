package observer

import "fmt"

type TrafficManagement struct {
	Observable
}

func (t *TrafficManagement) Notify(data any) {
	if p, ok := data.(PropertyChanged); ok {
		if p.Value.(int) >= 16 {
			fmt.Println("Congrats, you can drive now!")
			// we no longer care
			t.Unsubscribe(t)
		}
	}
}
