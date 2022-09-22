package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

const (
	k = 32
)

func main() {
	players := importJson()
	fmt.Println(players)
	// make two sample players
	playerOne := player{Name: "Nikola Jokic", Team: "Nuggets", elo: 1000, wins: 0, loses: 0}
	playerTwo := player{Name: "Joel Embiid", Team: "76ers", elo: 1000, wins: 0, loses: 0}
	// save their original scores
	oldOne := playerOne.elo
	oldTwo := playerTwo.elo
	// run it with player one winning
	playerOne, playerTwo = elo(playerOne, playerTwo, true)
	// output scores
	fmt.Printf("%s was ranked %d, but is now ranked %d\n", playerOne.Name, oldOne, playerOne.elo)
	fmt.Printf("%s was ranked %d, but is now ranked %d\n", playerTwo.Name, oldTwo, playerTwo.elo)

}

func importJson() []player {
	// Open our jsonFile
	jsonFile, err := os.Open("cmd/players.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var players Players

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &players)
	var allPlayers []player

	allPlayers = append(allPlayers, players.Players...)
	for i := range allPlayers {
		allPlayers[i].elo = 1000
	}

	return allPlayers
}

type Players struct {
	Players []player `json:"Players"`
}

type player struct {
	Name          string `json:"Name"`
	Team          string `json:"Team"`
	elo           int
	wins          int
	loses         int
	expectedScore float64
}

/*


    {
      "Name": "Hassan Whiteside",
      "Team": "Utah Jazz",
      "Position": "C",
      "Age": "32",
      "Height": "7' 0\"",
      "Height_i": "7.0",
      "Weight": "265",
      "College": "Marshall",
      "Salary": "1669178"
    }

Import players

Pick two players

Ask the person who's better

modify the scores

get rid of the loser

find someone else to rate against (with x score of the winner to keep it reasonble, 800?)

repeat

show ranking
*/

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
