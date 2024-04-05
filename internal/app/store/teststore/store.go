package teststore

import (
	"github.com/nizepart/rest-go/internal/app/store"
	"github.com/nizepart/rest-go/model"
)

type Store struct {
	userRepository          *UserRepository
	emailTemplateRepository *EmailTemplateRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepository
}

func (s *Store) EmailTemplate() store.EmailTemplateRepository {
	if s.emailTemplateRepository != nil {
		return s.emailTemplateRepository
	}

	s.emailTemplateRepository = &EmailTemplateRepository{
		store: s,
	}

	return s.emailTemplateRepository
}
