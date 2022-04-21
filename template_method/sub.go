package templateMethod

import "fmt"

func Sub() {
	templateMethod()
	fmt.Println()
	functionalTemplateMethod()
}

func templateMethod() {
	chess := NewGameOfChess()
	PlayGame(chess)
}

func functionalTemplateMethod() {
	turn, maxTurns, currentPlayer := 1, 10, 0

	start := func() {
		fmt.Println("Starting a game of chess.")
		fmt.Printf("Turn %d taken by player %d\n", turn, currentPlayer)
	}

	takeTurn := func() {
		turn++
		currentPlayer = (currentPlayer + 1) % 2
		fmt.Printf("Turn %d taken by player %d\n", turn, currentPlayer)
	}

	haveWinner := func() bool {
		return turn == maxTurns
	}

	winningPlayer := func() int {
		return currentPlayer
	}

	PlayGame2(start, takeTurn, haveWinner, winningPlayer)

}
