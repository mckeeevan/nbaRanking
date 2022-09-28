package elo

import (
	"testing"

	importjson "main.go/pkg/importJSON"
)

func TestUpdatePlayerOneWin(t *testing.T) {

	playerOne := importjson.Player{Name: "TestOne", Wins: 0, Loses: 0, Elo: 1000}
	playerTwo := importjson.Player{Name: "TestTwo", Wins: 0, Loses: 0, Elo: 1000}
	playerOne, playerTwo = winOdds(playerOne, playerTwo)
	gotPlayerOne, gotPlayerTwo := updatePlayerOneWin(playerOne, playerTwo)
	wantPlayerOne := importjson.Player{Name: "TestOne", Wins: 1, Loses: 0, Elo: 1064, ExpectedScore: .5}
	wantPlayerTwo := importjson.Player{Name: "TestTwo", Wins: 0, Loses: 1, Elo: 936, ExpectedScore: .5}

	if gotPlayerOne != wantPlayerOne {
		t.Errorf("Player One is Wrong")
		t.Error("Want Playerone", wantPlayerOne)
		t.Error("Got Playerone", gotPlayerOne)
		t.Error("Want Playertwo", wantPlayerTwo)
		t.Error("Got Playertwo", gotPlayerTwo)
	}

	if gotPlayerTwo != wantPlayerTwo {
		t.Errorf("Player Two is wrong")
		t.Error("Want Playerone", wantPlayerOne)
		t.Error("Got Playerone", gotPlayerOne)
		t.Error("Want Playertwo", wantPlayerTwo)
		t.Error("Got Playertwo", gotPlayerTwo)
	}

}

func TestUpdatePlayerTwoWin(t *testing.T) {

	playerOne := importjson.Player{Name: "TestOne", Wins: 0, Loses: 0, Elo: 1000}
	playerTwo := importjson.Player{Name: "TestTwo", Wins: 0, Loses: 0, Elo: 1000}
	playerOne, playerTwo = winOdds(playerOne, playerTwo)
	gotPlayerOne, gotPlayerTwo := updatePlayerTwoWin(playerOne, playerTwo)
	wantPlayerOne := importjson.Player{Name: "TestOne", Wins: 0, Loses: 1, Elo: 936, ExpectedScore: .5}
	wantPlayerTwo := importjson.Player{Name: "TestTwo", Wins: 1, Loses: 0, Elo: 1064, ExpectedScore: .5}

	if gotPlayerOne != wantPlayerOne {
		t.Errorf("Player One is Wrong")
		t.Error("Want Playerone", wantPlayerOne)
		t.Error("Got Playerone", gotPlayerOne)
		t.Error("Want Playertwo", wantPlayerTwo)
		t.Error("Got Playertwo", gotPlayerTwo)
	}

	if gotPlayerTwo != wantPlayerTwo {
		t.Errorf("Player Two is wrong")
		t.Error("Want Playerone", wantPlayerOne)
		t.Error("Got Playerone", gotPlayerOne)
		t.Error("Want Playertwo", wantPlayerTwo)
		t.Error("Got Playertwo", gotPlayerTwo)
	}

}

func TestEloPlayerOneWin(t *testing.T) {

	playerOne := importjson.Player{Name: "TestOne", Wins: 0, Loses: 0, Elo: 1000}
	playerTwo := importjson.Player{Name: "TestTwo", Wins: 0, Loses: 0, Elo: 1000}

	gotPlayerOne, gotPlayerTwo := Elo(playerOne, playerTwo, true)
	wantPlayerOne := importjson.Player{Name: "TestOne", Wins: 1, Loses: 0, Elo: 1064, ExpectedScore: .5}
	wantPlayerTwo := importjson.Player{Name: "TestTwo", Wins: 0, Loses: 1, Elo: 936, ExpectedScore: .5}

	if gotPlayerOne != wantPlayerOne {
		t.Errorf("Player One is Wrong")
		t.Error("Want Playerone", wantPlayerOne)
		t.Error("Got Playerone", gotPlayerOne)
		t.Error("Want Playertwo", wantPlayerTwo)
		t.Error("Got Playertwo", gotPlayerTwo)
	}

	if gotPlayerTwo != wantPlayerTwo {
		t.Errorf("Player Two is wrong")
		t.Error("Want Playerone", wantPlayerOne)
		t.Error("Got Playerone", gotPlayerOne)
		t.Error("Want Playertwo", wantPlayerTwo)
		t.Error("Got Playertwo", gotPlayerTwo)
	}

}

func TestEloPlayerTwoWin(t *testing.T) {

	playerOne := importjson.Player{Name: "TestOne", Wins: 0, Loses: 0, Elo: 1000}
	playerTwo := importjson.Player{Name: "TestTwo", Wins: 0, Loses: 0, Elo: 1000}
	gotPlayerOne, gotPlayerTwo := Elo(playerOne, playerTwo, false)
	wantPlayerOne := importjson.Player{Name: "TestOne", Wins: 0, Loses: 1, Elo: 936, ExpectedScore: .5}
	wantPlayerTwo := importjson.Player{Name: "TestTwo", Wins: 1, Loses: 0, Elo: 1064, ExpectedScore: .5}

	if gotPlayerOne != wantPlayerOne {
		t.Errorf("Player One is Wrong")
		t.Error("Want Playerone", wantPlayerOne)
		t.Error("Got Playerone", gotPlayerOne)
		t.Error("Want Playertwo", wantPlayerTwo)
		t.Error("Got Playertwo", gotPlayerTwo)
	}

	if gotPlayerTwo != wantPlayerTwo {
		t.Errorf("Player Two is wrong")
		t.Error("Want Playerone", wantPlayerOne)
		t.Error("Got Playerone", gotPlayerOne)
		t.Error("Want Playertwo", wantPlayerTwo)
		t.Error("Got Playertwo", gotPlayerTwo)
	}

}
