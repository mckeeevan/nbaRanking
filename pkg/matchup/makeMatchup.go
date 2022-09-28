package matchup

const (
	winsAllowed          = 3
	eloDifferenceAllowed = 400
)

func pickMatchup(data ModelData) (Matchup, int) {

	winInfo := winInfo{data.Winner, data.WinCount, data.WinChange}

	// check to see if there was a winner, and if the winner is allowed another match
	// are they over the max number of wins in a row

	if winCheck(winInfo) {
		// Add one person to winner
		two := getRandomPlayer(data)
		matchPlayers := PlayersMatchup{data.Players[data.Match.one], data.Players[two]}
		for validMatchCheck(matchPlayers) {
			two = getRandomPlayer(data)
			matchPlayers = PlayersMatchup{data.Players[data.Match.one], data.Players[two]}
		}
		data.Match.two = two
		return data.Match, data.WinCount + 1
	}

	one, two := getRandomPlayers(data)
	matchPlayers := PlayersMatchup{data.Players[one], data.Players[two]}

	for validMatchCheck(matchPlayers) {
		two = getRandomPlayer(data)
		matchPlayers = PlayersMatchup{data.Players[one], data.Players[two]}
	}

	return Matchup{one, two}, 0
}
