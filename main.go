package main

import (
	"CSVtoSQL/csv"
	"CSVtoSQL/sqlparser"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	tables, err := csv.CSVReader(args)
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	sqlFile, err := os.Create("statements.sql")
	for tn, table := range tables {
		tableName := args[tn][:strings.LastIndex(args[tn], ".")]
		sqlparser.Load(sqlFile, tableName, table)
	}
}
