package srp

import (
	"io/ioutil"
	"strings"
)

type Persistence struct {
	LineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entires, p.LineSeparator)), 0644)
}
