package sqlstore_test

import (
	"testing"

	"github.com/nizepart/mailbomber/internal/app/model"
	"github.com/nizepart/mailbomber/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestEmailTemplateRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("email_templates")
	s := sqlstore.New(db)

	et := model.TestEmailTemplate(t)
	assert.NoError(t, s.EmailTemplate().Create(et))
	assert.NotNil(t, et)
}

func TestEmailTemplateRepository_FindByID(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("email_templates")
	s := sqlstore.New(db)

	et := model.TestEmailTemplate(t)
	s.EmailTemplate().Create(et)
	etFound, err := s.EmailTemplate().FindByID(et.ID)
	assert.NoError(t, err)
	assert.Equal(t, et, etFound)
}
