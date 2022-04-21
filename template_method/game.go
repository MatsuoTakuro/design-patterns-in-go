package templateMethod

import "fmt"

type Game interface {
	Start()
	HaveWinner() bool
	TakeTurn()
	WinningPlayer() int
}

func PlayGame(g Game) {
	g.Start()
	for !g.HaveWinner() {
		g.TakeTurn()
	}
	fmt.Printf("Player %d wins.\n", g.WinningPlayer())
}

func PlayGame2(start, takeTurn func(),
	haveWinner func() bool,
	winningPlayer func() int) {
	start()
	for !haveWinner() {
		takeTurn()
	}
	fmt.Printf("Player %d wins.\n", winningPlayer())
}
