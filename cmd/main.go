package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	importjson "main.go/pkg/importJSON"
	"main.go/pkg/matchup"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	players := importjson.ImportJson()

	var players2 []importjson.Player

	// Put Nuggets players in a new slice
	for _, v := range players {
		if v.Team == "Denver Nuggets" {
			players2 = append(players2, v)
		}
	}

	// Make the players slice just Mavs and Nuggets players
	players = players2
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
}
