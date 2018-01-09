package common

import (
	"encoding/csv"
	"fmt"
	"os"
)

type DBKey struct {
	Username string
	Password string
}

func GetKeys(passwordsCSV string) (output map[string]DBKey) {
	var tempOutput = make(map[string]DBKey)
	file, err := os.Open(passwordsCSV)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for n, val := range rows {
		if n > 0 {
			var out DBKey
			out.Username = val[1]
			out.Password = val[2]
			tempOutput[val[0]] = out
		}
	}
	output = tempOutput
	return output
}
