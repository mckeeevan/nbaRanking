package elo

import (
	importjson "main.go/pkg/importJSON"
)

const (
	k = 32
)

func Elo(playerOne, playerTwo importjson.Player, playerOneWin bool) (importjson.Player, importjson.Player) {

	playerOne, playerTwo = winOdds(playerOne, playerTwo)

	if playerOneWin {
		// If player one wins move player ones ranking up and player two down
		playerOne.Elo = int(scoreChange(playerOne, true))
		playerTwo.Elo = int(scoreChange(playerTwo, false))
		playerOne.Wins++
		playerTwo.Loses++
	}
	if !playerOneWin {
		// If player two wins move player two's ranking up and player one down
		playerOne.Elo = int(scoreChange(playerOne, false))
		playerTwo.Elo = int(scoreChange(playerTwo, true))
		playerOne.Loses++
		playerTwo.Wins++
	}
	// Return both
	return playerOne, playerTwo
}
