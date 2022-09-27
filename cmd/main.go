package main

import (
	"math/rand"
	"time"

	"main.go/pkg/initialize"
	"main.go/pkg/model"
)
func main() {

	// make random seed
	rand.Seed(time.Now().UnixNano())

	// Import players for the first time
	players := initialize.Players(29.0, 10.0)

	// reopen ranked players
	// players := importjson.Ranked("cmd/scoredPlayers.json")

	model.Run(players, 5)

}

/*

web stuff
	basic UI
	Come up with better UI
	Make better UI

*/


