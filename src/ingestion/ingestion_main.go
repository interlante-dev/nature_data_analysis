package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// IngestionEngine is a struct that matches the ingestion_config.json
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
	ingestionEngines := GetIngestionEngines("/home/jint-dev/dev/nature_data_analysis/src/ingestion/ingestion_config.json")
	for _, val := range ingestionEngines.Engines {
		fmt.Println(val)
	}
	for {
	}
}

// GetIngestionEngines reads the ingestion_config.json
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
