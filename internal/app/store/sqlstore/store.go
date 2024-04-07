package sqlstore

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/nizepart/rest-go/internal/app/store"
)

type Store struct {
	db                      *sql.DB
	userRepository          *UserRepository
	emailTemplateRepository *EmailTemplateRepository
	emailScheduleRepository *EmailScheduleRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
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

func (s *Store) EmailSchedule() store.EmailScheduleRepository {
	if s.emailScheduleRepository != nil {
		return s.emailScheduleRepository
	}

	s.emailScheduleRepository = &EmailScheduleRepository{
		store: s,
	}

	return s.emailScheduleRepository
}
