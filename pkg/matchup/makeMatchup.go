package matchup

import (
	"fmt"
	"math"
	"math/rand"

	"main.go/pkg/elo"
	importjson "main.go/pkg/importJSON"
)

func Random(players []importjson.Player) []importjson.Player {

	one := rand.Intn(len(players))
	two := rand.Intn(len(players))
	for two == one || math.Abs(float64(players[one].Elo-players[two].Elo)) > 400 {
		two = rand.Intn(len(players))
	}

	fmt.Println(players[one])
	fmt.Println(players[two])
	var pick int
	fmt.Print("Player 1 or 2: ")
	fmt.Scanln(&pick)

	if pick == 1 {
		players[one], players[two] = elo.Elo(players[one], players[two], true)
		fmt.Println("You chose", players[one].Name, "as being better than", players[two].Name)
		fmt.Println(players[one].Name, "is now ranked", players[one].Elo)
		fmt.Println(players[two].Name, "is now ranked", players[two].Elo)
		// fmt.Println(players)
	}
	if pick == 2 {
		players[one], players[two] = elo.Elo(players[one], players[two], false)
		fmt.Println("You chose", players[two].Name, "as being better than", players[one].Name)
		fmt.Println(players[one].Name, "is now ranked", players[one].Elo)
		fmt.Println(players[two].Name, "is now ranked", players[two].Elo)
		// fmt.Println(players)
	}

	return players
}