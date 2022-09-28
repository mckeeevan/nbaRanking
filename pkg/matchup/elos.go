package matchup

import (
	"math"

	"main.go/pkg/elo"
	importjson "main.go/pkg/importJSON"
)

func updatePlayerElos(data ModelData, playerOneWin bool) ModelData {
	data.Players[data.Match.one], data.Players[data.Match.two] = changePlayerElos(data, data.Match.one, data.Match.two, playerOneWin)
	return data
}

func changePlayerElos(data ModelData, playerOne, playerTwo int, playerOneWin bool) (importjson.Player, importjson.Player) {
	return elo.Elo(data.Players[playerOne], data.Players[playerTwo], playerOneWin)
}

func differenceInElo(matchPlayers PlayersMatchup) float64 {
	return math.Abs(float64(matchPlayers.playerOne.Elo - matchPlayers.playerTwo.Elo))
}
