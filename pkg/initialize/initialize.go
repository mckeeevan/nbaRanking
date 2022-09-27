package initialize

import (
	importjson "main.go/pkg/importJSON"
)

func Players(minutes, games float64) []importjson.Player {

	// import player and season data
	players, seasonData := importjson.InitialData()
	// attach the stats to the players from N number of seasons
	players = combinePlayersAndStats(players, seasonData)
	// return the players, but only those who meet
	// a certain number of games and minutes played.
	return minuteAndGameRestricted(players, minutes, games)

}
