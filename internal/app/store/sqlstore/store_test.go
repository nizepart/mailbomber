package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=db user=postgres password=postgres dbname=mailbomber_test sslmode=disable"
	}

	os.Exit(m.Run())
}
