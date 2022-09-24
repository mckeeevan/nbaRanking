package matchup

import (
	"fmt"
	"math"
	"math/rand"

	"main.go/pkg/elo"
	importjson "main.go/pkg/importJSON"
)

const (
	winsAllowed = 3
)

type Matchup struct {
	one int
	two int
}

func pickMatchup(players []importjson.Player, match Matchup, winner bool, winCount int) (Matchup, int) {
	if winner && winCount < winsAllowed {
		// Add one person to winner
		two := rand.Intn(len(players))
		for match.one == two || math.Abs(float64(players[match.one].Elo-players[two].Elo)) > 400 {
			two = rand.Intn(len(players))
		}
		match.two = two
		winCount++
		return match, winCount
	}
	one := rand.Intn(len(players))
	two := rand.Intn(len(players))
	for two == one || math.Abs(float64(players[one].Elo-players[two].Elo)) > 400 {
		two = rand.Intn(len(players))
	}
	return Matchup{one, two}, 0
}

func Random(players []importjson.Player, matchIncoming Matchup, winner bool, winCount int, winChange bool) ([]importjson.Player, Matchup, bool, int, bool) {
	if winChange {
		winCount = 0
	}
	match, winCount := pickMatchup(players, matchIncoming, winner, winCount)

	displayMatchup(players, match)
	var pick int
	fmt.Print("Player 1 or 2: ")
	fmt.Scanln(&pick)

	if pick == 1 {
		players[match.one], players[match.two] = elo.Elo(players[match.one], players[match.two], true)
		fmt.Println("You chose", players[match.one].Name, "as being better than", players[match.two].Name)
		fmt.Println(players[match.one].Name, "is now ranked", players[match.one].Elo)
		fmt.Println(players[match.two].Name, "is now ranked", players[match.two].Elo)
		// fmt.Println(players)

		return players, match, true, winCount, false
	}
	if pick == 2 {
		players[match.one], players[match.two] = elo.Elo(players[match.one], players[match.two], false)
		fmt.Println("You chose", players[match.two].Name, "as being better than", players[match.one].Name)
		fmt.Println(players[match.one].Name, "is now ranked", players[match.one].Elo)
		fmt.Println(players[match.two].Name, "is now ranked", players[match.two].Elo)
		// fmt.Println(players)
		match.one = match.two

		return players, match, true, winCount, true
	}

	return players, match, false, winCount, true
}

func displayPlayerInfo(player importjson.Player) {
	fmt.Println(player.Name, "who in 2021-2022 averaged", "Points:", player.Points, "Assists:", player.Assists, "Rebounds:", player.Rebounds, "Blocks:", player.Blocks, "Games Played:", player.GP, "Games Started:", player.GS,
		"per game.")

}

func displayMatchup(players []importjson.Player, match Matchup) {

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	displayPlayerInfo(players[match.one])
	fmt.Println()
	displayPlayerInfo(players[match.two])
	fmt.Println()
	fmt.Println()
	fmt.Println()
}
