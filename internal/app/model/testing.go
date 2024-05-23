package model

import (
	"fmt"
	"testing"
	"time"

	"github.com/nizepart/mailbomber/internal/app"
)

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@example.org",
		Password: "password",
	}
}

func TestEmailTemplate(t *testing.T) *EmailTemplate {
	return &EmailTemplate{
		Subject:  "subject",
		Body:     fmt.Sprintf("Hello, %s!", "Andy"),
		BodyType: "text/plain",
	}
}

func TestEmailSchedule(t *testing.T) *EmailSchedule {
	location, _ := time.LoadLocation(app.GetValue("TZ", "Europe/Moscow").String())
	return &EmailSchedule{
		EmailTemplateID: 1,
		Recipients:      "email@example.org",
		ExecuteAfter:    time.Now().In(location).Add(3 * time.Hour),
		ExecutionPeriod: 1,
	}
}
