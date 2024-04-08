package model

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"strings"
	"time"
)

func requiredIf(cond bool) validation.RuleFunc {
	return func(value interface{}) error {
		if cond {
			return validation.Validate(value, validation.Required)
		}
		return nil
	}
}

func isTimestamp(t time.Time) validation.RuleFunc {
	return func(value interface{}) error {
		if _, ok := value.(time.Time); !ok {
			return errors.New("must be a valid timestamp")
		}
		if t.Before(time.Now()) {
			return errors.New("ExecuteAfter must be later than or equal to the current time")
		}
		return nil
	}
}

func areValidEmails(recipients string) validation.RuleFunc {
	return func(value interface{}) error {
		emailList := strings.Split(recipients, ",")
		for _, email := range emailList {
			email = strings.TrimSpace(email)
			if err := validation.Validate(email, is.Email); err != nil {
				return err
			}
		}
		return nil
	}
}
