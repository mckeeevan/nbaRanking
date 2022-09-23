package elo

import (
	"math"

	importjson "main.go/pkg/importJSON"
)

const (
	k = 32
)

func kRate(player importjson.Player) int {
	switch games := player.Wins + player.Loses; {
	case games < 5:
		return 128
	case games < 10:
		return 64
	case games < 15:
		return 45
	}
	return 32
}

func Elo(playerOne, playerTwo importjson.Player, playerOneWin bool) (importjson.Player, importjson.Player) {

	playerOne, playerTwo = winOdds(playerOne, playerTwo)

	if playerOneWin {
		// If player one wins move player ones ranking up and player two down
		playerOne.Elo = int(scoreChange(playerOne, playerOne.Elo, true))
		playerTwo.Elo = int(scoreChange(playerTwo, playerTwo.Elo, false))
		playerOne.Wins++
		playerTwo.Loses++
	}
	if !playerOneWin {
		// If player two wins move player two's ranking up and player one down
		playerOne.Elo = int(scoreChange(playerOne, playerOne.Elo, false))
		playerTwo.Elo = int(scoreChange(playerTwo, playerTwo.Elo, true))
		playerOne.Loses++
		playerTwo.Wins++
	}
	// Return both
	return playerOne, playerTwo
}

func winOdds(playerOne, playerTwo importjson.Player) (importjson.Player, importjson.Player) {
	// Update each players expected score
	playerOne.ExpectedScore = 1.0 / (1.0 + math.Pow(10, (float64(playerTwo.Elo)-float64(playerOne.Elo))/400.0))
	playerTwo.ExpectedScore = 1.0 / (1.0 + math.Pow(10, (float64(playerOne.Elo)-float64(playerTwo.Elo))/400.0))
	return playerOne, playerTwo
}

func scoreChange(player importjson.Player, rating int, win bool) float64 {
	if win {
		// if they win add to their score
		return float64(rating) + float64(kRate(player))*(1.0-player.ExpectedScore)
	}
	// if they lose subract from their score
	return float64(rating) + float64(kRate(player))*(0.0-player.ExpectedScore)
}
