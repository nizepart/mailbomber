package model

import validation "github.com/go-ozzo/ozzo-validation"

type EmailTemplate struct {
	ID      int    `json:"id"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func (et *EmailTemplate) Validate() error {
	return validation.ValidateStruct(
		et,
		validation.Field(&et.Subject, validation.Required, validation.Length(1, 100)),
		validation.Field(&et.Body, validation.Required),
	)
}
