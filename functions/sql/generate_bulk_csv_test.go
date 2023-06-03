package sql

import (
	"fmt"
	"testing"

	"github.com/LintaoAmons/toolbox/functions"
	"github.com/LintaoAmons/toolbox/functions/random"
)

type Person struct {
	Id         string
	Name       string
	Created_at string
}

func generateOneRow(id int) string {
	a := Person{
		Id:         fmt.Sprint(id),
		Name:       random.PickOneFromList("Lintao", "Amons").(string),
		Created_at: random.PostgresTimeString(),
	}
	return functions.StructToCsvRow(a, false)
}

func TestEfeedbackService_generateOneRow(t *testing.T) {
	result := generateOneRow(1)
	fmt.Println(result)
}

func Test_GenerateCsvParallel(t *testing.T) {
	for i := 0; i < 3; i++ {
		GenerateCSVParallel(i, 3, 10, RowGeneratorFunc(generateOneRow), "test")
	}
}
