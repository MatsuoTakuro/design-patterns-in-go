package observer

import (
	"fmt"
)

type DoctorService struct{}

func (d *DoctorService) Notify(data any) {
	fmt.Printf("A doctor has been called for %s", data.(string))
}
