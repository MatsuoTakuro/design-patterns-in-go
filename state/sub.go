package state

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Sub() {
	// classicImpl()
	handmadeState()
}

func classicImpl() {
	sw := NewSwitch()
	sw.On()
	sw.On()
	sw.Off()
	sw.Off()
}

func handmadeState() {
	state, exit := OnHook, OnHook
	for ok := true; ok; ok = (state != exit) {
		fmt.Println("The phone is currently", state)
		fmt.Println("Select a trigger:")

		for i := 0; i < len(rules[state]); i++ {
			rlt := rules[state][i]
			fmt.Println(strconv.Itoa(i), ":", rlt.Trigger)
		}

		input, _, _ := bufio.NewReader(os.Stdin).ReadLine()
		i, _ := strconv.Atoi(string(input))

		rlt := rules[state][i]
		state = rlt.hmState
	}
	fmt.Println("We are done using the phone")
}
