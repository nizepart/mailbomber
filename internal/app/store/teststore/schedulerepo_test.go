package teststore_test

import (
	"github.com/nizepart/rest-go/internal/app/store/teststore"
	"github.com/nizepart/rest-go/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmailScheduleRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.T(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}
