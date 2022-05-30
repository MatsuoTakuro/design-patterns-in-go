package mediator

import (
	"fmt"
	"strings"
)

type Member struct {
	Name       string
	Channel    *Channel
	messageLog []string
}

func NewMember(name string) *Member {
	return &Member{
		Name: name,
		Channel: &Channel{
			people: []*Member{},
		},
		messageLog: []string{},
	}
}

func (m *Member) Receive(sender, msg string) {
	s := fmt.Sprintf("%s: '%s'", sender, msg)
	fmt.Printf("[%s's chat session] %s\n", m.Name, s)
	m.messageLog = append(m.messageLog, s)
}

func (m *Member) SayHere(msg string) {
	m.Channel.Broadcast(m.Name, msg)
}

func (m *Member) DirectMessage(who, msg string) {
	m.Channel.Message(m.Name, who, msg)
}

func (m Member) String() string {
	var sb strings.Builder
	intro := fmt.Sprintf("I'm %s\n", m.Name)
	sb.WriteString(intro)
	sb.WriteString("Message logs;\n")
	for i, log := range m.messageLog {
		line := fmt.Sprintf(" %d: %s\n", i, log)
		sb.WriteString(line)
	}
	return sb.String()
}
