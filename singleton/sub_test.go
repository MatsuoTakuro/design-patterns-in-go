package singleton

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSub(t *testing.T) {
	cities := []string{"Seoul", "Mexico City"}
	tp := GetTotalPopulation(cities)
	ok := (tp == (17500000 + 17400000))
	msg := fmt.Sprintf("The total population of Seoul and Mexico City is %d", tp)
	assert.Equal(t, true, ok, msg)
}
