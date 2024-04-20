package model_test

import (
	"fmt"
	"github.com/nizepart/rest-go/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMessage_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		m       func() *model.Message
		isValid bool
	}{
		{
			name: "valid",
			m: func() *model.Message {
				return &model.Message{
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
