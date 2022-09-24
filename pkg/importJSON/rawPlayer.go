package importjson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func Import(file string) []Player {
	// Open our jsonFile
	jsonFile, err := os.Open(file)
	// if we os.Open returns an error then handle itß
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var players Players

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &players)

	// fmt.Println(players)

	var allPlayers []Player

	allPlayers = append(allPlayers, players.Players...)

	/*
		for i := range allPlayers {
			allPlayers[i].Elo = 1000
		}
	*/

	return allPlayers
}

func ImportSeasons(file string) []Season {
	// Open our jsonFile
	jsonFile, err := os.Open(file)
	// if we os.Open returns an error then handle itß
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var seasons Seasons

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &seasons)

	// fmt.Println(players)

	var allSeasons []Season

	allSeasons = append(allSeasons, seasons.Seasons...)

	/*
		for i := range allPlayers {
			allPlayers[i].Elo = 1000
		}
	*/

	return allSeasons
}
