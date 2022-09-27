package stats

import (
	importjson "main.go/pkg/importJSON"
)

func CombineSeasons(seasons []importjson.Season, numSeasons int) []importjson.Season {
	// make new variables
	var allSeasons, allPlayers []string

	// add the names of all players and seasons to slices
	for i := range seasons {
		allSeasons = append(allSeasons, seasons[i].Seasons)
		allPlayers = append(allPlayers, seasons[i].Player)
	}
	// remove duplicates
	allPlayers = removeDuplicateValues(allPlayers)
	allSeasons = removeDuplicateValues(allSeasons)
	// get N number of recent seasons
	mostRecentSeasons := mostRecentSeasons(allSeasons, numSeasons)
	// get rid of all the seasons that don't fit in N number of recent years
	seasons = reduceSeasons(seasons, mostRecentSeasons)
	//
	playerInfo := addAllPlayers(seasons, allPlayers)
	return playerInfo
}

func addAllPlayers(seasons []importjson.Season, allPlayers []string) []importjson.Season {
	var finished []importjson.Season
	var temp []importjson.Season
	for _, players := range allPlayers {
		for i, season := range seasons {
			if season.Player == players {
				temp = append(temp, seasons[i])
			}
		}
		finished = append(finished, combineSeasons(temp))
		temp = nil
	}
	return finished
}

func reduceSeasons(seasons []importjson.Season, years []string) []importjson.Season {
	var seasonsFromYears []importjson.Season
	for i, season := range seasons {
		for _, year := range years {
			if season.Seasons == year {
				seasonsFromYears = append(seasonsFromYears, seasons[i])
			}
		}
	}
	return seasonsFromYears
}

func mostRecentSeasons(seasons []string, n int) []string {
	var lastN []string
	for i := 0; i < n; i++ {
		if len(seasons)-i-1 == 0 {
			return lastN
		}
		lastN = append(lastN, seasons[len(seasons)-i-1])

	}
	return lastN
}
