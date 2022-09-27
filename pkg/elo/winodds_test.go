package elo

import (
	"testing"

	importjson "main.go/pkg/importJSON"
)

func TestWinOdds(t *testing.T) {

	playerOne := importjson.Player{Name: "Test", Wins: 0, Loses: 0, Elo: 1000}
	playerTwo := importjson.Player{Name: "Test", Wins: 0, Loses: 0, Elo: 1000}
	gotOne, gotTwo := winOdds(playerOne, playerTwo)
	want := .5

	if gotOne.ExpectedScore != want {
		t.Errorf("got %f, wanted %f", gotOne.ExpectedScore, want)
	}

	if gotTwo.ExpectedScore != want {
		t.Errorf("got %f, wanted %f", gotTwo.ExpectedScore, want)
	}

}
