package common

import (
	"database/sql"
	"fmt"
	"strings"
)

// ConnectToDB lets you connect to the mysql database (customize as needed)
func ConnectToDB(passwordFile string, databaseType string, databaseName string) (dbConn *sql.DB) {
	keysList := GetKeys(passwordFile)[databaseType]
	userpass := fmt.Sprintf("%s:%s@/%s", strings.TrimSpace(keysList.Username), strings.TrimSpace(keysList.Password), databaseName)
	db, err := sql.Open("mysql", userpass)
	if err != nil {
		fmt.Println(err)
	}
	return db
}

// ReadFromDB abstracts away some of the common database work
func ReadFromDB(query string, dbConn *sql.DB) (outputRows *sql.Rows, outputScanArgs []interface{}, outputValues []sql.RawBytes) {
	rows, err := dbConn.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	outputRows = rows
	outputScanArgs = scanArgs
	outputValues = values
	return outputRows, outputScanArgs, outputValues
}
