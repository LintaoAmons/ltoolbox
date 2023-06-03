package functions

import (
	"testing"
)

func Test_CreateTableToStruct(t *testing.T) {
	// Test case with a sample CREATE TABLE statement
	createTableStmt := `CREATE TABLE users (
		id integer,
		username varchar(100),
		email varchar(255),
		created_at timestamp with time zone,
		active boolean
	);`

	expectedStruct := `type TableName struct {
		id         int64
		username   string
		email      string
		created_at time.Time
		active     bool
	}`

	result, err := CreateTableToStruct(createTableStmt)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != expectedStruct {
		t.Errorf("Generated struct does not match expected result.\nExpected:\n%s\n\nGot:\n%s", expectedStruct, result)
	}
}

type Person struct {
	Name  string
	Age   int
	Email string
}

func Test_StructToCsvRow(t *testing.T) {
	person := Person{Name: "John Doe", Age: 30, Email: "johndoe@example.com"}
  StructToCsvRow(person, false)

	// result := StructToCsvRow(person, false)

	// if result != "John Doe,30,johndoe@example.com" {
	// 	t.Errorf("Generated struct does not match expected result.\nExpected:\n%s\n\nGot:\n%s", "John Doe,30,johndoe@example.com", result)
	// }
}
