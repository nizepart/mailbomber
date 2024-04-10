package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"time"
)

type EmailSchedule struct {
	ID              int       `json:"id"`
	EmailTemplateID int       `json:"email_template_id"`
	Recipients      string    `json:"recipients"`
	ExecuteAfter    time.Time `json:"execute_after"`
	ExecutionPeriod int       `json:"execution_period"`
}

func (es *EmailSchedule) Validate() error {
	return validation.ValidateStruct(
		es,
		validation.Field(&es.EmailTemplateID, validation.Required, validation.Min(1)),
		validation.Field(&es.Recipients, validation.Required, validation.By(areValidEmails())),
		validation.Field(&es.ExecuteAfter, validation.Required, validation.By(isValidTimestamp())),
		validation.Field(&es.ExecutionPeriod, validation.Min(0), validation.Max(168)),
	)
}
