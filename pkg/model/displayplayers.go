package model

import (
	"fmt"

	"main.go/pkg/matchup"
)

func displayRankedPlayers(data matchup.ModelData) {
	blankLines(4)
	for _, v := range data.Players {
		fmt.Printf("%25s has an elo of %6.0d\n", v.Name, v.Elo)
	}

}

func blankLines(lines int) {
	for i := 0; i < lines; i++ {
		fmt.Println()
	}
}
