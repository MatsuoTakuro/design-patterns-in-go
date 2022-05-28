package state

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var code string = "1234"
var state SwState = NewSwState()
var entry = strings.Builder{}

func ValidateEntryCode() {
	for {
		switch state {
		case Locked:
			// only reads input when you press Return
			r, _, _ := bufio.NewReader(os.Stdin).ReadRune()
			entry.WriteRune(r)

			if entry.String() == code {
				state = Unlocked
				break
			}

			if strings.Index(code, entry.String()) != 0 {
				state = Failed
			}

		case Failed:
			fmt.Println(state)
			entry.Reset()
			state = Locked

		case Unlocked:
			fmt.Println(state)
			return
		}
	}
}
