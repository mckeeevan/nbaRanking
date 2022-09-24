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
	GP            float64
	GS            float64
	Minutes       float64
	Rebounds      float64
	Points        float64
	Assists       float64
	Steals        float64
	Blocks        float64
}

type RawPlayers struct {
	RawPlayers []RawPlayer `json:"Players"`
}

type RawPlayer struct {
	Name          string `json:"namePlayer"`
	Team          string `json:"Team"`
	Elo           int
	Wins          int
	Loses         int
	ExpectedScore float64
}

type Seasons struct {
	Seasons []Season `json:"Seasons"`
}
type Season struct {
	Seasons  string  `json:"slugSeason"`
	Player   string  `json:"namePlayer"`
	GP       float64 `json:"gp"`
	GS       float64 `json:"gs"`
	Minutes  float64 `json:"minutesPerGame"`
	Rebounds float64 `json:"trebPerGame"`
	Points   float64 `json:"ptsPerGame"`
	Assists  float64 `json:"astPerGame"`
	Steals   float64 `json:"stlPerGame"`
	Blocks   float64 `json:"blkPerGame"`
}
