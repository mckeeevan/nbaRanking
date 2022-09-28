package matchup

import (
	"fmt"

	importjson "main.go/pkg/importJSON"
)

func displayPlayerInfo(player importjson.Player) {
	fmt.Printf("%25s averaged %5.2f Points %5.2f Assists %5.2f Rebounds %5.2f Blocks and %5.2f Steals over the last three seasons", player.Name, player.Points, player.Assists, player.Rebounds, player.Blocks, player.Steals)

}

func displayMatchup(players []importjson.Player, match Matchup) {

	blankLines(10)
	displayPlayerInfo(players[match.one])
	blankLines(1)
	displayPlayerInfo(players[match.two])
	blankLines(3)
}

func blankLines(lines int) {

	for i := 0; i < lines; i++ {
		fmt.Println()
	}
}
