package model

import (
	"sort"

	"main.go/pkg/matchup"
)

func sortPlayers(data matchup.ModelData) matchup.ModelData {
	// sort the slice
	sort.Slice(data.Players, func(i, j int) bool {
		return data.Players[i].Elo > data.Players[j].Elo
	})
	// print everyone
	return data
}
