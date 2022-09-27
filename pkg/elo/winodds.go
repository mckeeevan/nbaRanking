package elo

import (
	"math"

	importjson "main.go/pkg/importJSON"
)

func winOdds(playerOne, playerTwo importjson.Player) (importjson.Player, importjson.Player) {
	// Update each players expected score
	playerOne.ExpectedScore = 1.0 / (1.0 + math.Pow(10, (float64(playerTwo.Elo)-float64(playerOne.Elo))/400.0))
	playerTwo.ExpectedScore = 1.0 / (1.0 + math.Pow(10, (float64(playerOne.Elo)-float64(playerTwo.Elo))/400.0))
	return playerOne, playerTwo
}
