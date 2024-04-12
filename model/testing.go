package model

import (
	"fmt"
	"testing"
	"time"
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
	return &EmailSchedule{
		EmailTemplateID: 1,
		Recipients:      "email@example.org",
		ExecuteAfter:    time.Now().Add(5 * time.Minute),
		ExecutionPeriod: 1,
	}
}
