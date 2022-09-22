package main

import (
	"fmt"
	"math"
)

const (
	ratingChange = 32
)

func main() {
	fmt.Println("hey")

	jokic := player{name: "Nikola Jokic", team: "Nuggets", elo: 1000, wins: 0, loses: 0}
	embiid := player{name: "Joel Embiid", team: "76ers", elo: 1000, wins: 0, loses: 0}
	oldOne := jokic.elo
	oldTwo := embiid.elo
	jokic, embiid = elo(jokic, embiid, 1)
	fmt.Printf("Jokic was ranked %d, but is now ranked %d\n", oldOne, jokic.elo)
	fmt.Printf("Embiid was ranked %d, but is now ranked %d", oldTwo, embiid.elo)

}

type player struct {
	name          string
	team          string
	elo           int
	wins          int
	loses         int
	expectedScore float64
}

func elo(playerOne, playerTwo player, winner int) (player, player) {
	playerOne.expectedScore, playerTwo.expectedScore = winOdds(playerOne.elo, playerTwo.elo)
	if winner == 1 {
		fmt.Println("winner 1")
		fmt.Println(int(scoreChange(1, playerOne.expectedScore, playerOne.elo, true)))
		playerOne.elo = int(scoreChange(1, playerOne.expectedScore, playerOne.elo, true))
		playerTwo.elo = int(scoreChange(1, playerTwo.expectedScore, playerTwo.elo, false))
	}
	if winner == 2 {
		fmt.Println("winner 1")
		playerOne.elo = int(scoreChange(1, playerOne.expectedScore, playerOne.elo, false))
		playerTwo.elo = int(scoreChange(1, playerTwo.expectedScore, playerTwo.elo, true))

	}
	return playerOne, playerTwo

}

func winOdds(rB, rA int) (float64, float64) {
	rAExpectedScore := 1.0 / (1.0 + math.Pow(10, (float64(rB)-float64(rA))/400.0))
	rBExpectedScore := 1.0 / (1.0 + math.Pow(10, (float64(rA)-float64(rB))/400.0))
	return rAExpectedScore, rBExpectedScore
}

func scoreChange(score, expectedScore float64, rating int, win bool) float64 {
	if win == true {
		return float64(rating) + float64(ratingChange)*(score-expectedScore)
	}
	return float64(rating) - float64(ratingChange)*(score-expectedScore)

}
