package matchup

import "fmt"

func askWinner() int {
	var pick int
	fmt.Print("Player 1 or 2: ")
	fmt.Scanln(&pick)
	return pick
}
