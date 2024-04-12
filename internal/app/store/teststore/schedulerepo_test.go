package teststore_test

import (
	"github.com/nizepart/rest-go/internal/app/store/teststore"
	"github.com/nizepart/rest-go/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmailScheduleRepository_Create(t *testing.T) {
	s := teststore.New()
	es := model.TestEmailSchedule(t)
	assert.NoError(t, s.EmailSchedule().Create(es))
	assert.NotNil(t, es)
}
