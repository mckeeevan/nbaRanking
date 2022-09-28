package matchup

func playerOneWon(data ModelData) ModelData {

	// update Elos
	data = updatePlayerElos(data, true)
	return data
}

func playerTwoWon(data ModelData) ModelData {

	// update Elos
	data = updatePlayerElos(data, false)
	// make player one the previous winner
	data.Match.one = data.Match.two

	return data
}
