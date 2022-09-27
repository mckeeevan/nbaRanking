package initialize

import (
	importjson "main.go/pkg/importJSON"
	"main.go/pkg/stats"
)

func combinePlayersAndStats(players []importjson.Player, seasonData []importjson.Season) []importjson.Player {

	seasonData = stats.CombineSeasons(seasonData, 3)
	for i, j := range players {
		for _, v := range seasonData {
			if v.Player == j.Name {
				players = applyStats(players, v, i)

			}
			// give all players an initial Elo
			players[i].Elo = 1000
		}
	}
	return players
}

func applyStats(players []importjson.Player, v importjson.Season, i int) []importjson.Player {
	players[i].Points = v.Points
	players[i].Rebounds = v.Rebounds
	players[i].Assists = v.Assists
	players[i].Blocks = v.Blocks
	players[i].GP = v.GP
	players[i].GS = v.GS
	players[i].Steals = v.Steals
	players[i].Minutes = v.Minutes
	return players
}
