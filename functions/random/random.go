package random

import (
	"math/rand"
	"time"
)

func PostgresTimeString() string {
	// Set the seed for the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random datetime within a specific range
	min := time.Date(1990, time.April, 1, 0, 0, 0, 0, time.UTC)
	max := time.Date(2031, time.April, 30, 23, 59, 59, 999999999, time.UTC)
	randomTime := randomDateTime(min, max)

	// Format the datetime string
	return randomTime.Format("2006-01-02 15:04:05.999999999 -07:00")
}

func randomDateTime(min, max time.Time) time.Time {
	delta := max.Unix() - min.Unix()
	sec := rand.Int63n(delta + 1)
	return min.Add(time.Second * time.Duration(sec))
}

func PickOneFromList(i ...interface{}) interface{} {
	// Set the seed for the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random index within the range of the list
	randomIndex := rand.Intn(len(i))

	// Pick the random string from the list using the random index
	return i[randomIndex]
}
