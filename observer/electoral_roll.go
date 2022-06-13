package observer

import "fmt"

// electoral roll means 選挙人名簿 in japanese.
type ElectrocalRoll struct {
	Observable
}

func (e *ElectrocalRoll) Notify(data any) {
	if p, ok := data.(PropertyChanged); ok {
		if p.Name == "CanVote" && p.Value.(bool) {
			fmt.Println("Congrats, you can vote!")
			e.Unsubscribe(e)
		}
	}
}
