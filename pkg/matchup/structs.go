package matchup

import importjson "main.go/pkg/importJSON"

type ModelData struct {
	Players   []importjson.Player
	Match     Matchup
	Winner    bool
	WinCount  int
	WinChange bool
	// winInfo   winInfo
}

type Matchup struct {
	one int
	two int
}

type PlayersMatchup struct {
	playerOne importjson.Player
	playerTwo importjson.Player
}

type winInfo struct {
	winner bool
	count  int
	Change bool
}
