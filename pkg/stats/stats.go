package stats

import (
	"fmt"

	importjson "main.go/pkg/importJSON"
)

func CombineYears(seasons []importjson.Season) {
	var justLebron []importjson.Season
	for i, v := range seasons {
		if v.Player == "LeBron James" {
			justLebron = append(justLebron, seasons[i])
		}
	}
	var allSeasons []string

	fmt.Println(seasons[1].Seasons)
	for i, _ := range seasons {
		allSeasons = append(allSeasons, seasons[i].Seasons)
	}

	allSeasons = removeDuplicateValues(allSeasons)

	fmt.Println(allSeasons)
	// fmt.Println(justLebron)
	lastN := lastN(allSeasons, 3)
	fmt.Println(lastN)

	fmt.Println(checkSeasons(seasons, allSeasons))
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
