package email_test

import (
	"fmt"
	"github.com/nizepart/rest-go/internal/app/email"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMessage_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		m       func() *email.Message
		isValid bool
	}{
		{
			name: "valid",
			m: func() *email.Message {
				return &email.Message{
					From:     "from@example.com",
					To:       []string{"to@example.com"},
					Cc:       []string{"cc@example.com"},
					Subject:  "subject",
					Body:     fmt.Sprintf("Hello %s!", "Andy"),
					BodyType: "text/html",
					Attach:   "",
				}
			},
			isValid: true,
		},
		{
			name: "invalid from",
			m: func() *email.Message {
				m := &email.Message{
					From:     "",
					To:       []string{"to@example.com"},
					Cc:       []string{"cc@example.com"},
					Subject:  "subject",
					Body:     "body",
					BodyType: "text/html",
					Attach:   "",
				}
				return m
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.m().Validate())
			} else {
				assert.Error(t, tc.m().Validate())
			}
		})
	}
}
