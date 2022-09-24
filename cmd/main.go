package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"sort"
	"time"

	importjson "main.go/pkg/importJSON"
	"main.go/pkg/matchup"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	// reopen ranked players
	players := importjson.Ranked("cmd/scoredPlayers.json")

	// Import players for the first time
	// players := intialize(29.0, 10.0)

	run(players)

}

/*

Matchup winner stays (for X number)

combine season stats
	Attach these stats to players
web stuff
	basic UI
	Come up with better UI
	Make better UI


*/

func intialize(minutes, games float64) []importjson.Player {

	players := importjson.Import("cmd/playerlist.json")
	seasonData := importjson.ImportSeasons("cmd/playerstats.json")

	//players := importjson.Initialize("cmd/players.json")

	for i, j := range players {
		for _, v := range seasonData {
			if v.Player == j.Name && v.Seasons == "2021-22" {
				players[i].Points = v.Points
				players[i].Rebounds = v.Rebounds
				players[i].Assists = v.Assists
				players[i].Blocks = v.Blocks
				players[i].GP = v.GP
				players[i].GS = v.GS
				players[i].Steals = v.Steals
				players[i].Minutes = v.Minutes

			}
			players[i].Elo = 1000
		}

	}
	// Only have players with x minutes
	var minutesRestrictedPlayers []importjson.Player

	for i, v := range players {
		if v.Minutes > minutes && v.GP > games {
			minutesRestrictedPlayers = append(minutesRestrictedPlayers, players[i])
		}
	}
	// players = append([]importjson.Player(nil), players[:5]...)

	return minutesRestrictedPlayers

}

func run(players []importjson.Player) {
	players, match, winner := matchup.Random(players, matchup.Matchup{}, false)
	players, match, winner = matchup.Random(players, match, winner)
	players, match, winner = matchup.Random(players, match, winner)
	players, match, winner = matchup.Random(players, match, winner)
	players, match, winner = matchup.Random(players, match, winner)
	players, match, winner = matchup.Random(players, match, winner)
	players, match, winner = matchup.Random(players, match, winner)
	players, match, winner = matchup.Random(players, match, winner)
	players, match, winner = matchup.Random(players, match, winner)
	players, match, winner = matchup.Random(players, match, winner)
	players, match, winner = matchup.Random(players, match, winner)
	players, match, winner = matchup.Random(players, match, winner)
	players, match, winner = matchup.Random(players, match, winner)
	players, match, winner = matchup.Random(players, match, winner)
	players, match, winner = matchup.Random(players, match, winner)
	// sort the slice
	sort.Slice(players, func(i, j int) bool {
		return players[i].Elo > players[j].Elo
	})
	// print everyone

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	for _, v := range players {
		fmt.Println(v.Name, "has an Elo of", v.Elo)
	}

	var output importjson.Players
	output.Players = players
	jsonOutput, _ := json.Marshal(output)
	_ = ioutil.WriteFile("cmd/scoredPlayers.json", jsonOutput, 0644)
}
