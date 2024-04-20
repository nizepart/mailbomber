package model_test

import (
	"testing"
	"time"

	"github.com/nizepart/rest-go/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestEmailSchedule_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		es      func() *model.EmailSchedule
		isValid bool
	}{
		{
			name: "valid",
			es: func() *model.EmailSchedule {
				return &model.EmailSchedule{
					EmailTemplateID: 1,
					Recipients:      "test@example.org, test2@example.org",
					ExecuteAfter:    time.Now().Add(5 * time.Minute),
					ExecutionPeriod: 60,
				}
			},
			isValid: true,
		},
		{
			name: "valid with ExecutionPeriod 0 and single recipient",
			es: func() *model.EmailSchedule {
				return &model.EmailSchedule{
					EmailTemplateID: 1,
					Recipients:      "test@example.org",
					ExecuteAfter:    time.Now().Add(5 * time.Minute),
					ExecutionPeriod: 0,
				}
			},
			isValid: true,
		},
		{
			name: "invalid Recipients with invalid email address",
			es: func() *model.EmailSchedule {
				return &model.EmailSchedule{
					EmailTemplateID: 1,
					Recipients:      "invalid",
					ExecuteAfter:    time.Now().Add(5 * time.Minute),
					ExecutionPeriod: 0,
				}
			},
			isValid: false,
		},
		{
			name: "invalid empty Recipients",
			es: func() *model.EmailSchedule {
				return &model.EmailSchedule{
					EmailTemplateID: 1,
					Recipients:      "",
					ExecuteAfter:    time.Now().Add(5 * time.Minute),
					ExecutionPeriod: 0,
				}
			},
			isValid: false,
		},
		{
			name: "invalid ExecuteAfter",
			es: func() *model.EmailSchedule {
				return &model.EmailSchedule{
					EmailTemplateID: 1,
					Recipients:      "test@example.org",
					ExecuteAfter:    time.Now().Add(-time.Minute), // set ExecuteAfter to one minute before the current time
					ExecutionPeriod: 60,
				}
			},
			isValid: false,
		},
		{
			name: "invalid EmailTemplateID",
			es: func() *model.EmailSchedule {
				return &model.EmailSchedule{
					EmailTemplateID: 0,
					Recipients:      "test@example.org",
					ExecuteAfter:    time.Now().Add(5 * time.Minute),
					ExecutionPeriod: 0,
				}
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.es().Validate())
			} else {
				assert.Error(t, tc.es().Validate())
			}
		})
	}
}
