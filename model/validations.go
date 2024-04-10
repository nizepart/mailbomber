package model

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/nizepart/rest-go/internal/app"
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

func isValidTimestamp() validation.RuleFunc {
	return func(value interface{}) error {
		if valueTime, ok := value.(time.Time); ok {
			location, _ := time.LoadLocation(app.GetEnvString("TZ", "UTC"))
			if valueTime.Location().String() != location.String() {
				return errors.New("ExecuteAfter must be in server timezone")
			}
			if valueTime.Before(time.Now()) {
				return errors.New("ExecuteAfter must be later than or equal to the current time")
			}
		} else {
			return errors.New("Must be a valid timestamp")
		}
		return nil
	}
}

func areValidEmails() validation.RuleFunc {
	return func(value interface{}) error {
		emailList := strings.Split(value.(string), ",")
		for _, email := range emailList {
			email = strings.TrimSpace(email)
			if err := validation.Validate(email, is.Email); err != nil {
				return err
			}
		}
		return nil
	}
}
