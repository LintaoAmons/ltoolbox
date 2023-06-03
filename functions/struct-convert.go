package functions

import (
	"encoding/csv"
	"fmt"
	"reflect"
	"strings"
)

func CreateTableToStruct(createTableStmt string) (string, error) {
	lines := strings.Split(createTableStmt, "\n")
	fields := make([]string, 0)

	for _, line := range lines {
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, "(", " ")
		line = strings.ReplaceAll(line, ")", " ")
		line = strings.ReplaceAll(line, ";", " ")
		line = strings.TrimSpace(line)
		if line == "" || strings.Contains(line, "CREATE TABLE") {
			continue
		}

		fieldParts := strings.Fields(line)
		if len(fieldParts) < 2 {
			// fmt.Printf("invalid field line: %s", line)
			return "", fmt.Errorf("invalid field line: %s", line)
		}

		fieldName := fieldParts[0]
		fieldType := fieldParts[1]

		goField := fmt.Sprintf("%s %s", fieldName, getGoType(fieldType))
		fields = append(fields, goField)
	}

	structFields := strings.Join(fields, "\n\t")
	structDefinition := fmt.Sprintf("type TableName struct {\n\t%s\n}", structFields)

	return structDefinition, nil
}

func StructToCsvRow(s interface{}, withHeader bool) string {
	// Open the CSV file for writing
	var csvString strings.Builder

	// Create a CSV writer
	writer := csv.NewWriter(&csvString)

	// Get the type of the struct
	structType := reflect.TypeOf(s)

	// Write the header row
	if withHeader {
		headerRow := make([]string, structType.NumField())
		for i := 0; i < structType.NumField(); i++ {
			headerRow[i] = structType.Field(i).Name
		}
		writer.Write(headerRow)
	}
	// Write the data row
	dataRow := make([]string, structType.NumField())
	for i := 0; i < structType.NumField(); i++ {
		fieldValue := reflect.ValueOf(s).Field(i)
		dataRow[i] = fmt.Sprintf("%v", fieldValue.Interface())
	}
	writer.Write(dataRow)

	writer.Flush()
	// Check for any writer error
	if writer.Error() != nil {
		panic(writer.Error().Error())
	}

	// Get the CSV data as a string
	return csvString.String()
}

func getGoType(pgType string) string {
	switch pgType {
	case "integer", "bigint":
		return "int64"
	case "character varying", "text", "varchar":
		return "string"
	case "boolean":
		return "bool"
	case "real", "double precision":
		return "float64"
	case "date", "timestamp":
		return "time.Time"
	default:
		return "interface{}"
	}
}
