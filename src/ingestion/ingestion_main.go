package main

import (
	"encoding/json"
	"fmt"
	"ingestion/keys"
	"io/ioutil"
)

type IngestionEngine struct {
	Engines []struct {
		Engine        string   `json:"engine"`
		CurrentDate   string   `json:"current_as_of"`
		Frequency     int      `json:"frequency"`
		RequestString []string `json:"request_strings"`
	} `json:"engines"`
}

func main() {
	fmt.Println("main ingestion")
	// apiKeys := make(map[string][])
	keysList := keys.GetKeys("/home/jint-dev/passwords.csv")
	ingestionEngines := GetIngestionEngines("/home/jint-dev/dev/nature_data_analysis/src/ingestion/ingestion_config.json")
	for _, val := range ingestionEngines.Engines {
		fmt.Println(val)
	}
	fmt.Println(keysList)
	for {
	}
}

func GetIngestionEngines(filePath string) (output IngestionEngine) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	var engines IngestionEngine
	if err := json.Unmarshal(data, &engines); err != nil {
		fmt.Println(err)
	}
	output = engines
	return output
}
