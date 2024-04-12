package sqlstore_test

import (
	"github.com/nizepart/rest-go/internal/app"
	"github.com/nizepart/rest-go/internal/app/store/sqlstore"
	"github.com/nizepart/rest-go/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEmailScheduleRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("email_schedule")
	s := sqlstore.New(db)

	et := model.TestEmailTemplate(t)
	s.EmailTemplate().Create(et)
	es := model.TestEmailSchedule(t)
	es.EmailTemplateID = et.ID
	assert.NoError(t, s.EmailSchedule().Create(es))
	assert.NotNil(t, es)
}

func TestEmailScheduleRepository_SelectExecutables(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("email_schedule")
	s := sqlstore.New(db)

	et := model.TestEmailTemplate(t)
	s.EmailTemplate().Create(et)
	es := model.TestEmailSchedule(t)
	es.EmailTemplateID = et.ID
	location, _ := time.LoadLocation(app.GetEnvString("TZ", "UTC"))
	es.ExecuteAfter = time.Now().In(location).Add(-6 * time.Minute)
	s.EmailSchedule().Create(es)

	_, err := s.EmailSchedule().SelectExecutables()
	assert.NoError(t, err)
	assert.NotNil(t, es)
}
