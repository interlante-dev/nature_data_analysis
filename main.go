package main

import (
	"common"
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

// APIKeys is a struct that matches the APIKey stored in the mysql db
type APIKeys struct {
	KeyID    int64  `json:"key_id"`
	KeyValue string `json:"key_value"`
	Email    string `json:"email"`
	Source   string `json:"source"`
}

// GetAPIKeys does just that, it gets the api keys from the mysql db
func GetAPIKeys(query string, dbConn *sql.DB) (output []APIKeys) {
	rows, scanArgs, values := common.ReadFromDB(query, dbConn)
	outputAPIKeys := []APIKeys{}
	for rows.Next() {
		var apiKey APIKeys
		err := rows.Scan(scanArgs...)
		if err != nil {
			fmt.Println(err)
		}
		KeyID, err := strconv.ParseInt(string(values[0]), 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		apiKey.KeyID = KeyID
		apiKey.KeyValue = string(values[1])
		apiKey.Email = string(values[4])
		apiKey.Source = string(values[5])
		outputAPIKeys = append(outputAPIKeys, apiKey)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}
	output = outputAPIKeys
	return output
}

func main() {
	databaseType := "mysql"
	databaseName := "NATURE_DATA_ANALYSIS"
	passwordFile := "/home/jint-dev/passwords.csv"
	db := common.ConnectToDB(passwordFile, databaseType, databaseName)
	defer db.Close()

	apiOutput := GetAPIKeys("SELECT * FROM secret_keys", db)
	fmt.Println(apiOutput)
	fmt.Println(apiOutput[0].Source)

}
