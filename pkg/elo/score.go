package elo

import (
	importjson "main.go/pkg/importJSON"
)

func Elo(playerOne, playerTwo importjson.Player, playerOneWin bool) (importjson.Player, importjson.Player) {
	// check to make sure the players have a valid Elo score
	playerOne = checkForElo(playerOne)
	playerTwo = checkForElo(playerTwo)
	if playerOneWin {
		playerOne, playerTwo = updatePlayerOneWin(playerOne, playerTwo)
	}
	if !playerOneWin {
		playerOne, playerTwo = updatePlayerTwoWin(playerOne, playerTwo)
	}
	// Return both
	return playerOne, playerTwo
}

func checkForElo(player importjson.Player) importjson.Player {
	if player.Elo == 0 && player.Wins+player.Loses == 0 {
		player.Elo = 1000
		return player
	}
	if player.Elo >= 0 {
		player.Elo = 1
		return player
	}
	return player
}

func updatePlayerOneWin(playerOne, playerTwo importjson.Player) (importjson.Player, importjson.Player) {

	// If player one wins move playerOne's ranking up and playerTwo down
	playerOne, playerTwo = winOdds(playerOne, playerTwo)
	playerOne.Elo = int(scoreChange(playerOne, true))
	playerTwo.Elo = int(scoreChange(playerTwo, false))
	playerOne.Wins++
	playerTwo.Loses++
	return playerOne, playerTwo
}

func updatePlayerTwoWin(playerOne, playerTwo importjson.Player) (importjson.Player, importjson.Player) {

	// If player two wins move player two's ranking up and player one down
	playerOne, playerTwo = winOdds(playerOne, playerTwo)
	playerOne.Elo = int(scoreChange(playerOne, false))
	playerTwo.Elo = int(scoreChange(playerTwo, true))
	playerOne.Loses++
	playerTwo.Wins++
	return playerOne, playerTwo
}
