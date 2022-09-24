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
	/*
		data := importjson.Import("cmd/playerlist.json")
		fmt.Println(data[1])
	*/

	seasonsData := importjson.ImportSeasons("cmd/playerstats.json")

	stuff(seasonsData)

}

func stuff(seasonData []importjson.Season) {

	rand.Seed(time.Now().UnixNano())
	// players := importjson.Ranked("cmd/scoredPlayers.json")
	players := importjson.Initialize("cmd/players.json")

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
		}

	}
	run(players)
	// players = append([]importjson.Player(nil), players[:5]...)

}

func run(players []importjson.Player) {
	players = matchup.Random(players)
	players = matchup.Random(players)
	players = matchup.Random(players)
	players = matchup.Random(players)
	players = matchup.Random(players)
	players = matchup.Random(players)
	players = matchup.Random(players)
	players = matchup.Random(players)
	players = matchup.Random(players)
	players = matchup.Random(players)
	players = matchup.Random(players)
	players = matchup.Random(players)
	players = matchup.Random(players)
	players = matchup.Random(players)
	players = matchup.Random(players)
	players = matchup.Random(players)
	players = matchup.Random(players)
	players = matchup.Random(players)
	players = matchup.Random(players)
	players = matchup.Random(players)
	players = matchup.Random(players)
	players = matchup.Random(players)
	players = matchup.Random(players)
	players = matchup.Random(players)
	players = matchup.Random(players)
	players = matchup.Random(players)
	players = matchup.Random(players)
	players = matchup.Random(players)
	players = matchup.Random(players)
	players = matchup.Random(players)
	players = matchup.Random(players)
	players = matchup.Random(players)
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
