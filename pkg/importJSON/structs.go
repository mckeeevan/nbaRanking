package importjson

type Players struct {
	Players []Player `json:"Players"`
}

type Player struct {
	Name          string `json:"Name"`
	Team          string `json:"Team"`
	Elo           int
	Wins          int
	Loses         int
	ExpectedScore float64
}
