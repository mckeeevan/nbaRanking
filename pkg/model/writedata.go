package model

import (
	"encoding/json"
	"io/ioutil"

	importjson "main.go/pkg/importJSON"
	"main.go/pkg/matchup"
)

func writeData(data matchup.ModelData, fileName string) {
	// made a Players object
	var output importjson.Players
	// put all of the players into this object
	output.Players = data.Players
	// Marshal this data into a json
	jsonOutput, _ := json.Marshal(output)
	// Write the json into a file
	_ = ioutil.WriteFile(fileName, jsonOutput, 0644)
}
