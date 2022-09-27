package elo

import importjson "main.go/pkg/importJSON"

func scoreChange(player importjson.Player, rating int, win bool) float64 {
	if win {
		// if they win add to their score
		return float64(rating) + float64(kRate(player))*(1.0-player.ExpectedScore)
	}
	// if they lose subract from their score
	return float64(rating) + float64(kRate(player))*(0.0-player.ExpectedScore)
}
