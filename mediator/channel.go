package mediator

import (
	"fmt"
	"strings"
)

type Channel struct {
	people []*Member
}

func (c *Channel) Broadcast(src, msg string) {
	for _, m := range c.people {
		if m.Name != src {
			m.Receive(src, msg)
		}
	}
}

func (c *Channel) Join(m *Member) {
	joinMsg := fmt.Sprintf("%s joins the channel", m.Name)
	c.Broadcast("Room", joinMsg)

	m.Channel = c
	c.people = append(c.people, m)
}

func (c *Channel) Message(src, dst, msg string) {
	for _, m := range c.people {
		if m.Name == dst {
			m.Receive(src, msg)
		}
	}
}

func (c Channel) String() string {
	var sb strings.Builder
	sb.WriteString("This channel's members;\n")
	for i, m := range c.people {
		line := fmt.Sprintf(" %d,member: %s\n", i, m.Name)
		sb.WriteString(line)
	}
	return sb.String()
}
