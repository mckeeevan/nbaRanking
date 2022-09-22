package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"time"
)

const (
	k = 32
)

func main() {

	rand.Seed(time.Now().UnixNano())
	players := importJson()

	one := rand.Intn(len(players))
	two := rand.Intn(len(players))

	if two == one {
		two = rand.Intn(len(players))
	}

	fmt.Println(players[one])
	fmt.Println(players[two])

	var i int

	fmt.Print("Player 1 or 2: ")
	fmt.Scanln(&i)
	if i == 1 {
		players[one], players[two] = elo(players[one], players[two], true)
		fmt.Println("You chose", players[one].Name, "as being better than", players[two].Name)
		fmt.Println(players[one].Name, "is now ranked", players[one].elo)
		fmt.Println(players[two].Name, "is now ranked", players[two].elo)
	}
	if i == 2 {
		players[one], players[two] = elo(players[one], players[two], false)
		fmt.Println("You chose", players[two].Name, "as being better than", players[one].Name)
		fmt.Println(players[one].Name, "is now ranked", players[one].elo)
		fmt.Println(players[two].Name, "is now ranked", players[two].elo)
	}

	players[one], players[two] = elo(players[one], players[two], true)

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
	return float64(rating) - float64(k)*(score-expectedScore)ß
}
