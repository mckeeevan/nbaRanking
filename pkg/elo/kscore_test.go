package elo

import (
	"testing"

	importjson "main.go/pkg/importJSON"
)

func TestKRate(t *testing.T) {

	player := importjson.Player{Name: "Test", Wins: 0, Loses: 0}
	got := kRate(player)
	want := 128

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	player.Wins = 10
	got = kRate(player)
	want = 64

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	player.Wins = 20
	got = kRate(player)
	want = 45

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	player.Wins = 30
	got = kRate(player)
	want = 32

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
