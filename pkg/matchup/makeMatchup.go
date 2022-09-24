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

func pickMatchup(data ModelData) (Matchup, int) {
	if data.Winner && data.WinCount < winsAllowed {
		// Add one person to winner
		two := rand.Intn(len(data.Players))
		for data.Match.one == two || math.Abs(float64(data.Players[data.Match.one].Elo-data.Players[two].Elo)) > 400 {
			two = rand.Intn(len(data.Players))
		}
		data.Match.two = two
		data.WinCount++
		return data.Match, data.WinCount
	}
	one := rand.Intn(len(data.Players))
	two := rand.Intn(len(data.Players))
	for two == one || math.Abs(float64(data.Players[one].Elo-data.Players[two].Elo)) > 400 {
		two = rand.Intn(len(data.Players))
	}
	return Matchup{one, two}, 0
}

type ModelData struct {
	Players   []importjson.Player
	Match     Matchup
	Winner    bool
	WinCount  int
	WinChange bool
}

func Random(data ModelData) ModelData {
	if data.WinChange {
		data.WinCount = 0
	}
	match, winCount := pickMatchup(data)

	displayMatchup(data.Players, match)
	var pick int
	fmt.Print("Player 1 or 2: ")
	fmt.Scanln(&pick)

	if pick == 1 {
		data.Players[match.one], data.Players[match.two] = elo.Elo(data.Players[match.one], data.Players[match.two], true)
		fmt.Println("You chose", data.Players[match.one].Name, "as being better than", data.Players[match.two].Name)
		fmt.Println(data.Players[match.one].Name, "is now ranked", data.Players[match.one].Elo)
		fmt.Println(data.Players[match.two].Name, "is now ranked", data.Players[match.two].Elo)
		// fmt.Println(players)

		return ModelData{data.Players, match, true, winCount, false}
	}
	if pick == 2 {
		data.Players[match.one], data.Players[match.two] = elo.Elo(data.Players[match.one], data.Players[match.two], false)
		fmt.Println("You chose", data.Players[match.two].Name, "as being better than", data.Players[match.one].Name)
		fmt.Println(data.Players[match.one].Name, "is now ranked", data.Players[match.one].Elo)
		fmt.Println(data.Players[match.two].Name, "is now ranked", data.Players[match.two].Elo)
		// fmt.Println(players)data.
		match.one = match.two

		return ModelData{data.Players, match, true, winCount, true}
	}

	return ModelData{data.Players, match, false, winCount, true}
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
