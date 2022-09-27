package elo

import (
	importjson "main.go/pkg/importJSON"
)

func kRate(player importjson.Player) int {
	switch games := player.Wins + player.Loses; {
	case games < 10:
		return 128
	case games < 20:
		return 64
	case games < 30:
		return 45
	}
	return 32
}

func add(x, y int) int {

	return x + y
}
