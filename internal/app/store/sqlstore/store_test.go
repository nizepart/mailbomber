package sqlstore_test

import (
	"os"
	"testing"

	"github.com/nizepart/rest-go/internal/app"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = app.GetValue("DB_TEST_URL", "host=db user=postgres password=postgres dbname=mailbomber_test sslmode=disable").String()
	os.Exit(m.Run())
}
