package singleton

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type DummyDatabase struct {
	dummyData map[string]int
}

func (d *DummyDatabase) GetPopulation(name string) int {
	if len(d.dummyData) == 0 {
		d.dummyData = map[string]int{
			"alpha": 1,
			"beta":  2,
			"gamma": 3}
	}
	return d.dummyData[name]
}

func TestSub(t *testing.T) {
	cities := []string{"Seoul", "Mexico City"}
	tp := GetTotalPopulation(cities)
	ok := (tp == (17500000 + 17400000))
	msg := fmt.Sprintf("The total population of Seoul and Mexico City is %d", tp)
	assert.Equal(t, true, ok, msg)

	names := []string{"alpha", "gamma"}
	tp = GetTotalPopulationEx(&DummyDatabase{}, names)
	ok = tp == 4
	msg = fmt.Sprintf("The total population of alpha and gamma is %d", tp)
	assert.Equal(t, true, ok, msg)
}
