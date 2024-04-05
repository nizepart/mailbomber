package sqlstore

import "github.com/nizepart/rest-go/model"

type EmailTemplateRepository struct {
	store *Store
}

func (r *EmailTemplateRepository) Create(et *model.EmailTemplate) error {
	if err := et.Validate(); err != nil {
		return err
	}

	return r.store.db.QueryRow("INSERT INTO email_templates (subject, body) VALUES ($1, $2) RETURNING id", et.Subject, et.Body).Scan(&et.ID)
}
