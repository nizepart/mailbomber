package teststore_test

import (
	"github.com/nizepart/rest-go/internal/app"
	"github.com/nizepart/rest-go/internal/app/store/teststore"
	"github.com/nizepart/rest-go/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEmailScheduleRepository_Create(t *testing.T) {
	s := teststore.New()

	et := model.TestEmailTemplate(t)
	s.EmailTemplate().Create(et)
	es := model.TestEmailSchedule(t)
	es.EmailTemplateID = et.ID
	assert.NoError(t, s.EmailSchedule().Create(es))
	assert.NotNil(t, es)
}

// TODO: I don't know how to check the SelectExecutables function because of model validation
func TestEmailScheduleRepository_SelectExecutables(t *testing.T) {
	s := teststore.New()

	et := model.TestEmailTemplate(t)
	s.EmailTemplate().Create(et)
	es := model.TestEmailSchedule(t)
	es.EmailTemplateID = et.ID
	location, _ := time.LoadLocation(app.GetValue("TZ", "Europe/Moscow").String())
	es.ExecuteAfter = time.Now().In(location).Add(5 * time.Minute)
	errCreateSchedule := s.EmailSchedule().Create(es)
	assert.NoError(t, errCreateSchedule)
	_, err := s.EmailSchedule().SelectExecutables()
	assert.NoError(t, err)
}

func TestEmailScheduleRepository_UpdateExecutionTime(t *testing.T) {
	s := teststore.New()

	et := model.TestEmailTemplate(t)
	s.EmailTemplate().Create(et)
	es := model.TestEmailSchedule(t)
	es.EmailTemplateID = et.ID
	location, _ := time.LoadLocation(app.GetValue("TZ", "Europe/Moscow").String())
	es.ExecuteAfter = time.Now().In(location).Add(5 * time.Minute)
	errCreateSchedule := s.EmailSchedule().Create(es)
	assert.NoError(t, errCreateSchedule)
	timeBefore := es.ExecuteAfter
	err := s.EmailSchedule().UpdateExecutionTime(es)
	assert.NoError(t, err)
	assert.Greater(t, es.ExecuteAfter, timeBefore)
}
