package elo

import (
	"testing"

	importjson "main.go/pkg/importJSON"
)

func TestScoreChange(t *testing.T) {

	playerOne := importjson.Player{Name: "Test", Wins: 0, Loses: 0, Elo: 1000, ExpectedScore: .5}
	playerOneNewScore := scoreChange(playerOne, true)
	want := 1064.0
	if playerOneNewScore != want {
		t.Errorf("got %f, wanted %f", playerOneNewScore, want)
	}

	playerOneNewScore = scoreChange(playerOne, false)
	want = 936
	if playerOneNewScore != want {
		t.Errorf("got %f, wanted %f", playerOneNewScore, want)
	}
}
