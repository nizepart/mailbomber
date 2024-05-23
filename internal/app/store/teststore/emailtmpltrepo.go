package teststore

import (
	"github.com/nizepart/mailbomber/internal/app/model"
	"github.com/nizepart/mailbomber/internal/app/store"
)

type EmailTemplateRepository struct {
	store          *Store
	emailTemplates map[int]*model.EmailTemplate
}

func (r *EmailTemplateRepository) Create(et *model.EmailTemplate) error {
	if err := et.Validate(); err != nil {
		return err
	}

	et.ID = len(r.emailTemplates) + 1
	r.emailTemplates[et.ID] = et

	return nil
}

func (r *EmailTemplateRepository) FindByID(id int) (*model.EmailTemplate, error) {
	et, ok := r.emailTemplates[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return et, nil
}
