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
	// players := importjson.Ranked("cmd/scoredPlayers.json")
	players := importjson.Initialize("cmd/players.json")

	// players = append([]importjson.Player(nil), players[:5]...)

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
