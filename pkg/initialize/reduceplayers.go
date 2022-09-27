package initialize

import importjson "main.go/pkg/importJSON"

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
