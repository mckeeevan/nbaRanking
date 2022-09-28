package matchup

import "math/rand"

func Random(data ModelData) ModelData {
	if data.WinChange {
		data.WinCount = 0
	}
	data.Match, data.WinCount = pickMatchup(data)
	displayMatchup(data.Players, data.Match)
	pick := askWinner()

	return handlePicks(data, pick)
}

func getRandomPlayer(data ModelData) int {
	return rand.Intn(len(data.Players))
}

func getRandomPlayers(data ModelData) (int, int) {
	return rand.Intn(len(data.Players)), rand.Intn(len(data.Players))
}
