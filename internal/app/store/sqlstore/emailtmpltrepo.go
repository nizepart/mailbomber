package sqlstore

import (
	"database/sql"
	"github.com/nizepart/rest-go/internal/app/model"
	"github.com/nizepart/rest-go/internal/app/store"
)

type EmailTemplateRepository struct {
	store *Store
}

func (r *EmailTemplateRepository) Create(et *model.EmailTemplate) error {
	if err := et.Validate(); err != nil {
		return err
	}

	return r.store.db.QueryRow("INSERT INTO email_templates (subject, body, body_type) VALUES ($1, $2, $3) RETURNING id", et.Subject, et.Body, et.BodyType).Scan(&et.ID)
}

func (r *EmailTemplateRepository) FindByID(id int) (*model.EmailTemplate, error) {
	et := &model.EmailTemplate{}
	if err := r.store.db.QueryRow("SELECT id, subject, body, body_type FROM email_templates where id = $1", id).Scan(&et.ID, &et.Subject, &et.Body, &et.BodyType); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return et, nil
}
