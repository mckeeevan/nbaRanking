package matchup

import (
	"fmt"
	"math"
	"math/rand"

	"main.go/pkg/elo"
	importjson "main.go/pkg/importJSON"
)

const (
	winsAllowed          = 3
	eloDifferenceAllowed = 400
)

type Matchup struct {
	one int
	two int
}

type PlayersMatchup struct {
	playerOne importjson.Player
	playerTwo importjson.Player
}

func winCheck(winInfo winInfo) bool {
	if winInfo.winner && winInfo.count < winsAllowed {
		return true
	}
	return false
}

func pickMatchup(data ModelData) (Matchup, int) {

	// chek to see if there was a winner, and if the winner is allowed another match
	// are they over the max number of wins in a row
	if data.Winner && data.WinCount < winsAllowed {
		// Add one person to winner
		two := getRandomPlayer(data)
		matchPlayers := PlayersMatchup{data.Players[data.Match.one], data.Players[two]}
		for validMatchCheck(matchPlayers) {
			two = getRandomPlayer(data)
			matchPlayers = PlayersMatchup{data.Players[data.Match.one], data.Players[two]}
		}
		data.Match.two = two
		return data.Match, data.WinCount + 1
	}

	one, two := getRandomPlayers(data)
	matchPlayers := PlayersMatchup{data.Players[one], data.Players[two]}

	for validMatchCheck(matchPlayers) {
		two = getRandomPlayer(data)
		matchPlayers = PlayersMatchup{data.Players[one], data.Players[two]}
	}

	return Matchup{one, two}, 0
}

func validMatchCheck(matchPlayers PlayersMatchup) bool {
	if matchPlayers.playerOne == matchPlayers.playerTwo || differenceInElo(matchPlayers) > eloDifferenceAllowed {
		return true
	}
	return false
}

func differenceInElo(matchPlayers PlayersMatchup) float64 {
	return math.Abs(float64(matchPlayers.playerOne.Elo - matchPlayers.playerTwo.Elo))
}

func getRandomPlayer(data ModelData) int {
	return rand.Intn(len(data.Players))
}

func getRandomPlayers(data ModelData) (int, int) {
	return rand.Intn(len(data.Players)), rand.Intn(len(data.Players))
}

type ModelData struct {
	Players   []importjson.Player
	Match     Matchup
	Winner    bool
	WinCount  int
	WinChange bool
	// winInfo   winInfo
}

func Random(data ModelData) ModelData {
	if data.WinChange {
		data.WinCount = 0
	}
	data.Match, data.WinCount = pickMatchup(data)
	displayMatchup(data.Players, data.Match)
	pick := askWinner()

	return handlePicks(data, pick)
}

type winInfo struct {
	winner bool
	count  int
	Change bool
}

func handlePicks(data ModelData, pick int) ModelData {

	if pick == 1 {
		data = updatePlayerElos(data, true)
		return ModelData{data.Players, data.Match, true, data.WinCount, false}
	}
	if pick == 2 {
		data = playerTwoWon(data)
		return ModelData{data.Players, data.Match, true, data.WinCount, true}
	}

	return ModelData{data.Players, data.Match, false, data.WinCount, true}
}

func playerTwoWon(data ModelData) ModelData {

	data = updatePlayerElos(data, false)
	// make player one the previous winner
	data.Match.one = data.Match.two

	return data
}

func askWinner() int {
	var pick int
	fmt.Print("Player 1 or 2: ")
	fmt.Scanln(&pick)
	return pick
}

func updatePlayerElos(data ModelData, playerOneWin bool) ModelData {
	data.Players[data.Match.one], data.Players[data.Match.two] = changePlayerElos(data, data.Match.one, data.Match.two, playerOneWin)
	return data
}

func changePlayerElos(data ModelData, playerOne, playerTwo int, playerOneWin bool) (importjson.Player, importjson.Player) {
	return elo.Elo(data.Players[playerOne], data.Players[playerTwo], playerOneWin)
}

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
