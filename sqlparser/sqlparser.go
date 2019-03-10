package sqlparser

import (
	"fmt"
	"os"
	"regexp"
)

func Format(input string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	return reg.ReplaceAllString(input, "")
}

func Load(file *os.File, tableName string, cols [][]string) {
	columnNames := cols[0]
	myType := make(map[int]string)
	var colNameString string
	colType := []string{"VARCHAR(50)", "INTEGER", "FLOAT", "DATETIME"}
	file.WriteString("CREATE TABLE IF NOT EXISTS " + Format(tableName) + " (\n")
	for i, col := range columnNames {
		var input int
		fmt.Println("Define the type of " + col + " in " + tableName + ": \n1:VARCHAR, 2:INT, 3:FLOAT, 4:DATETIME")
		fmt.Scanf("%d", &input)
		for colType[input] == "" {
			fmt.Println("Invalid type of " + col + " in " + tableName + ", please enter again: \n1:VARCHAR, 2:INT, 3:DOUBLE, 4:DATE, 5:DATETIME, 6:TIMESTAMP")
			fmt.Scanf("%d", &input)
		}
		file.WriteString(Format(col) + " " + colType[input-1] + "")
		colNameString = colNameString + Format(col)
		myType[i] = colType[input-1]

		if i != len(columnNames)-1 {
			file.WriteString(",\n")
			colNameString = colNameString + ", "
		}
	}
	file.WriteString("\n);\n")

	insertLine := "INSERT INTO " + Format(tableName) + " (" + colNameString + ") \nVALUES("
	for _, col := range cols[1:] {
		file.WriteString(insertLine)
		for idx, val := range col {
			switch myType[idx] {
			case "VARCHAR(50)":
				file.WriteString("\"" + val + "\"")
			case "DATETIME":
				file.WriteString("datetime(" + val + ")")
			default:
				file.WriteString(val)
			}
			if idx != len(col)-1 {
				file.WriteString(", ")
			}
		}

		file.WriteString(");\n")
	}
	return
}
