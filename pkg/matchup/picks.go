package matchup

func handlePicks(data ModelData, pick int) ModelData {

	if pick == 1 {
		data = playerOneWon(data)
		return ModelData{data.Players, data.Match, true, data.WinCount, false}
	}
	if pick == 2 {
		data = playerTwoWon(data)
		return ModelData{data.Players, data.Match, true, data.WinCount, true}
	}

	return ModelData{data.Players, data.Match, false, data.WinCount, true}
}
