package matchup

func validMatchCheck(matchPlayers PlayersMatchup) bool {
	if matchPlayers.playerOne == matchPlayers.playerTwo || differenceInElo(matchPlayers) > eloDifferenceAllowed {
		// Return that it's not a valid matchup
		return true
	}
	// Return that it's a valid matchup
	return false
}

func winCheck(winInfo winInfo) bool {
	if winInfo.winner && winInfo.count < winsAllowed {
		return true
	}
	return false
}
