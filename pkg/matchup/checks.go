package matchup

func validMatchCheck(matchPlayers PlayersMatchup) bool {
	if matchPlayers.playerOne == matchPlayers.playerTwo || differenceInElo(matchPlayers) > eloDifferenceAllowed {
		return true
	}
	return false
}

func winCheck(winInfo winInfo) bool {
	if winInfo.winner && winInfo.count < winsAllowed {
		return true
	}
	return false
}
