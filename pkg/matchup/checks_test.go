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
		expect      bool
	}{
		{"Winner, first win", true, 0, true, true},
		{"Winner, 3rd win", true, 4, false, false},
	}
	for _, v := range testData {

		got := winCheck(winInfo{winner: v.inputWinner, count: v.inputCount, Change: v.inputChange})
		if got != v.expect {
			t.Errorf("%s got %t, wanted %t", v.description, got, v.expect)
		}
	}

}

func TestValidMatchCheck(t *testing.T) {

	testData := []struct {
		description    string
		inputPlayerOne importjson.Player
		inputPlayerTwo importjson.Player
		expect         bool
	}{
		{"Same Player", importjson.Player{Name: "One"}, importjson.Player{Name: "One"}, true},
		{"Different Players", importjson.Player{Name: "One"}, importjson.Player{Name: "Two"}, false},
		{"Different Players, two far apart elo", importjson.Player{Name: "One", Elo: 1000}, importjson.Player{Name: "One", Elo: 2000}, true},
	}
	for _, v := range testData {

		got := validMatchCheck(PlayersMatchup{v.inputPlayerOne, v.inputPlayerTwo})
		if got != v.expect {
			t.Errorf("%s got %t, wanted %t", v.description, got, v.expect)
		}
	}

}
