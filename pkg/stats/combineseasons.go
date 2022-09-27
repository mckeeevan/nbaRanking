package stats

import importjson "main.go/pkg/importJSON"

func combineSeasons(seasons []importjson.Season) importjson.Season {
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
	// Average the stats over the number of games played
	combined.Assists = combined.Assists / combined.GP
	combined.Points = combined.Points / combined.GP
	combined.Rebounds = combined.Rebounds / combined.GP
	combined.Minutes = combined.Minutes / combined.GP
	combined.Blocks = combined.Blocks / combined.GP
	combined.Steals = combined.Steals / combined.GP
	return combined
}
