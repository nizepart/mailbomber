package teststore_test

import (
	"github.com/nizepart/rest-go/internal/app/store/teststore"
	"github.com/nizepart/rest-go/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmailTemplateRepository_Create(t *testing.T) {
	s := teststore.New()

	et := model.TestEmailTemplate(t)
	assert.NoError(t, s.EmailTemplate().Create(et))
	assert.NotNil(t, et)
}

func TestEmailTemplateRepository_FindByID(t *testing.T) {
	s := teststore.New()

	et := model.TestEmailTemplate(t)
	s.EmailTemplate().Create(et)
	etFound, err := s.EmailTemplate().FindByID(et.ID)
	assert.NoError(t, err)
	assert.Equal(t, et, etFound)
}
