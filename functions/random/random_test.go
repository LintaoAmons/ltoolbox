package random

import (
	"fmt"
	"testing"
)

func TestPostgresTimeString(t *testing.T) {
	fmt.Println(PostgresTimeString())
}

func Test_PickOneFromList(t *testing.T) {
  result := PickOneFromList("aa", "bb", "cc")
	fmt.Println(result)
}
