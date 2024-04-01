package store_test

import (
	"github.com/nizepart/rest-go/internal/app/store"
	"github.com/nizepart/rest-go/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "user@example.org"
	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)

	u, err := s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, u.Email, email)
}
