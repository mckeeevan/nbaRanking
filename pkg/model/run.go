package model

import (
	importjson "main.go/pkg/importJSON"
	"main.go/pkg/matchup"
)

func Run(players []importjson.Player, runs int) {

	// Make first matchup
	data := matchup.ModelData{Players: players, Match: matchup.Matchup{},
		Winner: false, WinCount: 0, WinChange: true}
	// run the model the number of times it's requested
	data = modelLoop(data, runs)
	// sort players by Elo
	data = sortPlayers(data)
	// display a list of all the players
	displayRankedPlayers(data)
	// Output data
	writeData(data, "cmd/scoredPlayers.json")
}

func modelLoop(data matchup.ModelData, runs int) matchup.ModelData {

	// How many times to run
	for i := 0; i < runs; i++ {
		// run the model
		data = matchup.Random(data)
	}

	return data
}
