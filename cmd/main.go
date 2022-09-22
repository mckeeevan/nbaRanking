package main

import (
	"fmt"
	"math"
)

const (
	k = 32
)

func main() {
	// make two sample players
	playerOne := player{name: "Nikola Jokic", team: "Nuggets", elo: 1000, wins: 0, loses: 0}
	playerTwo := player{name: "Joel Embiid", team: "76ers", elo: 1000, wins: 0, loses: 0}
	// save their original scores
	oldOne := playerOne.elo
	oldTwo := playerTwo.elo
	// run it with player one winning
	playerOne, playerTwo = elo(playerOne, playerTwo, true)
	// output scores
	fmt.Printf("%s was ranked %d, but is now ranked %d\n", playerOne.name, oldOne, playerOne.elo)
	fmt.Printf("%s was ranked %d, but is now ranked %d\n", playerTwo.name, oldTwo, playerTwo.elo)

}

type player struct {
	name          string
	team          string
	elo           int
	wins          int
	loses         int
	expectedScore float64
}

func elo(playerOne, playerTwo player, playerOneWin bool) (player, player) {
	playerOne.expectedScore, playerTwo.expectedScore = winOdds(playerOne.elo, playerTwo.elo)
	if playerOneWin {
		// If player one wins move player ones ranking up and player two down
		playerOne.elo = int(scoreChange(1, playerOne.expectedScore, playerOne.elo, true))
		playerTwo.elo = int(scoreChange(1, playerTwo.expectedScore, playerTwo.elo, false))
	}
	if !playerOneWin {
		// If player two wins move player two's ranking up and player one down
		playerOne.elo = int(scoreChange(1, playerOne.expectedScore, playerOne.elo, false))
		playerTwo.elo = int(scoreChange(1, playerTwo.expectedScore, playerTwo.elo, true))
	}
	// Return both
	return playerOne, playerTwo
}

func winOdds(rB, rA int) (float64, float64) {
	// Update each players expected score
	rAExpectedScore := 1.0 / (1.0 + math.Pow(10, (float64(rB)-float64(rA))/400.0))
	rBExpectedScore := 1.0 / (1.0 + math.Pow(10, (float64(rA)-float64(rB))/400.0))
	return rAExpectedScore, rBExpectedScore
}

func scoreChange(score, expectedScore float64, rating int, win bool) float64 {
	if win {
		// if they win add to their score
		return float64(rating) + float64(k)*(score-expectedScore)
	}
	// if they lose subract from their score
	return float64(rating) - float64(k)*(score-expectedScore)

}
