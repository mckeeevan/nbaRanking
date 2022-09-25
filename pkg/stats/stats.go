package stats

import (
	"fmt"

	importjson "main.go/pkg/importJSON"
)

func CombineYears(seasons []importjson.Season) []importjson.Season {
	var justLebron []importjson.Season
	for i, v := range seasons {
		if v.Player == "LeBron James" {
			justLebron = append(justLebron, seasons[i])
		}
	}
	var allSeasons, allPlayers []string

	fmt.Println(seasons[1].Seasons)
	for i, _ := range seasons {
		allSeasons = append(allSeasons, seasons[i].Seasons)
		allPlayers = append(allPlayers, seasons[i].Player)
	}
	allPlayers = removeDuplicateValues(allPlayers)
	allSeasons = removeDuplicateValues(allSeasons)
	lastN := lastN(allSeasons, 3)
	seasons = checkSeasons(seasons, lastN)
	playerInfo := addAllPlayers(seasons, allPlayers)
	fmt.Println(playerInfo)

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
		finished = append(finished, addSeasons(temp))
		temp = nil
	}

	return finished
}

func addSeasons(seasons []importjson.Season) importjson.Season {
	var combined importjson.Season
	for _, v := range seasons {
		combined.GP = combined.GP + v.GP
		combined.GS = combined.GS + v.GS
		combined.Assists = combined.Assists + (v.Assists * v.GP)
		combined.Points = combined.Points + (v.Points * v.GP)
		combined.Rebounds = combined.Rebounds + (v.Rebounds * v.GP)
		combined.Minutes = combined.Minutes + (v.Minutes * v.GP)
		combined.Blocks = combined.Blocks + (v.Blocks * v.GP)
		combined.Steals = combined.Steals + (v.Steals * v.GP)
		combined.Player = v.Player
	}

	combined.Assists = combined.Assists / combined.GP
	combined.Points = combined.Points / combined.GP
	combined.Rebounds = combined.Rebounds / combined.GP
	combined.Minutes = combined.Minutes / combined.GP
	combined.Blocks = combined.Blocks / combined.GP
	combined.Steals = combined.Steals / combined.GP
	return combined
}

func checkSeasons(seasons []importjson.Season, years []string) []importjson.Season {
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

func removeDuplicateValues(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func lastN(seasons []string, n int) []string {
	var lastN []string
	for i := 0; i < n; i++ {
		if len(seasons)-i-1 == 0 {
			return lastN
		}
		lastN = append(lastN, seasons[len(seasons)-i-1])

	}
	return lastN
}

func CombinePlayerTeams(players []importjson.Player) {

}
