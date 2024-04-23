package store

import (
	"github.com/nizepart/mailbomber/internal/app/model"
)

type UserRepository interface {
	Create(*model.User) error
	FindByID(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}

type EmailTemplateRepository interface {
	Create(*model.EmailTemplate) error
	FindByID(int) (*model.EmailTemplate, error)
}

type EmailScheduleRepository interface {
	Create(*model.EmailSchedule) error
	SelectExecutables() ([]*model.EmailSchedule, error)
	UpdateExecutionTime(*model.EmailSchedule) error
	Delete(es *model.EmailSchedule) error
}
