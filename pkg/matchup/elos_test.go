package matchup

import (
	"testing"

	importjson "main.go/pkg/importJSON"
)

func TestDifferenceInElo(t *testing.T) {

	testData := []struct {
		description    string
		inputPlayerOne importjson.Player
		inputPlayerTwo importjson.Player
		expected       float64
	}{
		{"Same Elo", importjson.Player{Name: "One", Elo: 1000}, importjson.Player{Name: "Two", Elo: 1000}, 0.0},
		{"Player one higher Elo", importjson.Player{Name: "One", Elo: 2000}, importjson.Player{Name: "Two", Elo: 1000}, 1000.0},
		{"Player two higher Elo", importjson.Player{Name: "One", Elo: 1000}, importjson.Player{Name: "Two", Elo: 2000}, 1000.},
	}
	for _, testCase := range testData {

		t.Run(testCase.description, func(t *testing.T) {
			winCheckData := PlayersMatchup{testCase.inputPlayerOne, testCase.inputPlayerTwo}
			if result := differenceInElo(winCheckData); result != testCase.expected {
				t.Errorf("%s got %f, want %f", testCase.description, result, testCase.expected)
			}
		})
	}

}
