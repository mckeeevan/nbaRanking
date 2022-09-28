package elo

import (
	"testing"

	importjson "main.go/pkg/importJSON"
)

func TestCheckForElo(t *testing.T) {

	player := importjson.Player{Name: "Test", Wins: 0, Loses: 0}

	testData := []struct {
		description string
		inputElo    int
		inputGames  int
		expect      int
	}{
		{"Score of 0 with no games", 0, 0, 1000},
		{"Score of 0 with games played", 0, 1, 1},
		{"Score of -100 with no games", -100, 0, 1000},
		{"Score of -100 with games played", -100, 1, 1},
		{"Score of 500", 500, 1, 500},
		{"Score of 1000", 1000, 1, 1000},
	}

	for _, v := range testData {
		player.Elo = v.inputElo
		player.Loses = v.inputGames
		got := checkForElo(player)
		if got.Elo != v.expect {
			t.Errorf("input of %s elo got %d, wanted %d", v.description, got.Elo, v.expect)
		}
	}

}

func TestUpdatePlayerOneWin(t *testing.T) {

	playerOne := importjson.Player{Name: "TestOne", Wins: 0, Loses: 0, Elo: 1000}
	playerTwo := importjson.Player{Name: "TestTwo", Wins: 0, Loses: 0, Elo: 1000}

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
