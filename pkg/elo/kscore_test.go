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
}
