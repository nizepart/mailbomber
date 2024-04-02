package store

import "github.com/nizepart/rest-go/model"

type UserRepository interface {
	Create(*model.User) error
	FindByID(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}
