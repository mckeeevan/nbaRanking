package initialize

import (
	importjson "main.go/pkg/importJSON"
)

func Players(minutes, games float64) []importjson.Player {

	players, seasonData := importjson.InitialData()

	players = combinePlayersAndStats(players, seasonData)

	return minuteAndGameRestricted(players, minutes, games)

}
