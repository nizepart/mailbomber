package teststore

import (
	"github.com/nizepart/rest-go/model"
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
