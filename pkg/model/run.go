package model


import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"

	importjson "main.go/pkg/importJSON"
	"main.go/pkg/matchup"
)



func Run(players []importjson.Player, runs int) {

	// Make first matchup
	initial := matchup.ModelData{Players: players, Match: matchup.Matchup{}, Winner: false, WinCount: 0, WinChange: true}
	data := matchup.Random(initial)

	// How many times to run
	for i := 0; i < runs; i++ {
		data = matchup.Random(data)
	}

	data = sortPlayers(data)
	displayRankedPlayers(data)

	// Output data
	// writeData(data, "cmd/scoredPlayers.json")
}

func writeData(data matchup.ModelData, fileName string) {
	var output importjson.Players
	output.Players = data.Players
	jsonOutput, _ := json.Marshal(output)
	_ = ioutil.WriteFile(fileName, jsonOutput, 0644)

}

func sortPlayers(data matchup.ModelData) matchup.ModelData {
	// sort the slice
	sort.Slice(data.Players, func(i, j int) bool {
		return data.Players[i].Elo > data.Players[j].Elo
	})
	// print everyone
	return data
}

func displayRankedPlayers(data matchup.ModelData) {
	blankLines(4)
	for _, v := range data.Players {
		fmt.Printf("%25s has an elo of %6.0d\n", v.Name, v.Elo)
	}

}

func blankLines(lines int) {
	for i := 0; i < lines; i++ {
		fmt.Println()
	}
}