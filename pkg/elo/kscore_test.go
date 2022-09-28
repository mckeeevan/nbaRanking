package elo

import (
	"testing"

	importjson "main.go/pkg/importJSON"
)

func TestKRate(t *testing.T) {

	testData := []struct {
		description string
		input       int
		expect      int
	}{
		{"0 Games", 0, 128},
		{"10 Games", 10, 64},
		{"20 Games", 20, 45},
		{"30 Games", 30, 32},
		{"50 Games", 50, 32},
	}

	player := importjson.Player{Name: "Test", Wins: 0, Loses: 0}
	for i, v := range testData {
		player.Loses = testData[i].input
		got := kRate(player)
		if got != v.expect {
			t.Errorf("input %s got %d, wanted %d", v.description, got, v.expect)
		}
	}
}
