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

const (
	runs = 50
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// players := importjson.Import("cmd/playerlist.json")
	// seasonData := importjson.ImportSeasons("cmd/playerstats.json")

	// Import players for the first time
	// 	players := intialize(29.0, 10.0)
	// stats.CombineYears(seasonData)

	// reopen ranked players
	players := importjson.Ranked("cmd/scoredPlayers.json")

	run(players)

}

/*

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

	// Make first matchup
	initial := matchup.ModelData{Players: players, Match: matchup.Matchup{}, Winner: false, WinCount: 0, WinChange: true}
	data := matchup.Random(initial)

	// How many times to run
	for i := 0; i < runs; i++ {
		data = matchup.Random(data)
	}
	// sort the slice
	sort.Slice(players, func(i, j int) bool {
		return players[i].Elo > players[j].Elo
	})
	// print everyone

	blankLines(4)
	for _, v := range data.Players {
		fmt.Println(v.Name, "has an Elo of", v.Elo)
	}

	var output importjson.Players
	output.Players = data.Players
	jsonOutput, _ := json.Marshal(output)
	_ = ioutil.WriteFile("cmd/scoredPlayers.json", jsonOutput, 0644)
}

func blankLines(lines int) {
	for i := 0; i < lines; i++ {
		fmt.Println()

	}
}
