package matchup

import (
	"testing"

	importjson "main.go/pkg/importJSON"
)

func TestWinCheck(t *testing.T) {

	testData := []struct {
		description string
		inputWinner bool
		inputCount  int
		inputChange bool
		expected    bool
	}{
		{"Winner, first win", true, 0, true, true},
		{"Winner, 3rd win", true, 4, false, false},
	}
	for _, testCase := range testData {

		t.Run(testCase.description, func(t *testing.T) {
			winCheckData := winInfo{winner: testCase.inputWinner, count: testCase.inputCount, Change: testCase.inputChange}
			if result := winCheck(winCheckData); result != testCase.expected {
				t.Errorf("%s got %t, wanted %t", testCase.description, result, testCase.expected)
			}
		})
	}

}

func TestValidMatchCheck(t *testing.T) {

	testData := []struct {
		description    string
		inputPlayerOne importjson.Player
		inputPlayerTwo importjson.Player
		expected       bool
	}{
		{"Same Player", importjson.Player{Name: "One"}, importjson.Player{Name: "One"}, true},
		{"Different Players", importjson.Player{Name: "One"}, importjson.Player{Name: "Two"}, false},
		{"Different Players, too far apart elo", importjson.Player{Name: "One", Elo: 1000}, importjson.Player{Name: "One", Elo: 2000}, true},
		{"Same Player, too far apart elo", importjson.Player{Name: "One", Elo: 1000}, importjson.Player{Name: "One", Elo: 2000}, true},
	}
	for _, testCase := range testData {
		t.Run(testCase.description, func(t *testing.T) {
			if result := validMatchCheck(PlayersMatchup{testCase.inputPlayerOne, testCase.inputPlayerTwo}); result != testCase.expected {
				t.Errorf("%s got %t, wanted %t", testCase.description, result, testCase.expected)
			}
		})
	}

}
