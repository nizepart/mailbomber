package sqlstore_test

import (
	"github.com/nizepart/rest-go/internal/app"
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = app.GetEnvString("DB_TEST_URL", "host=db user=postgres password=postgres dbname=mailbomber_test sslmode=disable")
	os.Exit(m.Run())
}
