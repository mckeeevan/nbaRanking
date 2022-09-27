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
	players := initialize.Players(18.0, 25.0)

	// reopen ranked players
	// players := importjson.Reimport("cmd/scoredPlayers.json")

	model.Run(players, 1)

}

/*
web stuff
	basic UI
	Come up with better UI
	Make better UI
*/
