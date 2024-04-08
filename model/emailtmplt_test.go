package model_test

import (
	"fmt"
	"github.com/nizepart/rest-go/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmailTemplate_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		et      func() *model.EmailTemplate
		isValid bool
	}{
		{
			name: "valid",
			et: func() *model.EmailTemplate {
				return &model.EmailTemplate{
					Subject:  "subject",
					Body:     fmt.Sprintf("Hello %s!", "Andy"),
					BodyType: "text/plain",
				}
			},
			isValid: true,
		},
		{
			name: "no body type",
			et: func() *model.EmailTemplate {
				return &model.EmailTemplate{
					Subject:  "subject",
					Body:     fmt.Sprintf("Hello %s!", "Andy"),
					BodyType: "",
				}
			},
			isValid: false,
		},
		{
			name: "invalid body type",
			et: func() *model.EmailTemplate {
				return &model.EmailTemplate{
					Subject:  "subject",
					Body:     fmt.Sprintf("Hello %s!", "Andy"),
					BodyType: "test/test",
				}
			},
			isValid: false,
		},
		{
			name: "invalid subject",
			et: func() *model.EmailTemplate {
				return &model.EmailTemplate{
					Subject:  "",
					Body:     fmt.Sprintf("Hello %s!", "Andy"),
					BodyType: "text/plain",
				}
			},
			isValid: false,
		},
		{
			name: "invalid body",
			et: func() *model.EmailTemplate {
				return &model.EmailTemplate{
					Subject:  "subject",
					Body:     "",
					BodyType: "text/html",
				}
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.et().Validate())
			} else {
				assert.Error(t, tc.et().Validate())
			}
		})
	}
}
