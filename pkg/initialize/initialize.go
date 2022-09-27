package initialize

import (

	importjson "main.go/pkg/importJSON"
)


func getInitialData() ([]importjson.Player, []importjson.Season) {
	players := importjson.Import("cmd/playerlist.json")
	seasonData := importjson.ImportSeasons("cmd/playerstats.json")

	return players, seasonData
}


func combinePlayersAndStats(players []importjson.Player, seasonData []importjson.Season) []importjson.Player {

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

func minuteAndGameRestricted(players []importjson.Player, minutes, games float64) []importjson.Player {

	// Only have players with x minutes
	var minutesRestrictedPlayers []importjson.Player

	// check for players with two few minutes and games
	for i, v := range players {
		if v.Minutes > minutes && v.GP > games {
			minutesRestrictedPlayers = append(minutesRestrictedPlayers, players[i])
		}
	}

	return minutesRestrictedPlayers
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