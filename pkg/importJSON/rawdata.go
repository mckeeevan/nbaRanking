package importjson

func InitialData() ([]Player, []Season) {
	players := NBAPlayers("cmd/playerlist.json")
	seasonData := SeasonsData("cmd/playerstats.json")
	return players, seasonData
}