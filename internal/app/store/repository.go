package store

import "github.com/nizepart/rest-go/model"

type UserRepository interface {
	Create(*model.User) error
	FindByID(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}

type EmailTemplateRepository interface {
	Create(*model.EmailTemplate) error
	FindByID(int) (*model.EmailTemplate, error)
}
