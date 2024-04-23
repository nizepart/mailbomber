package teststore

import (
	"fmt"
	"time"

	"github.com/nizepart/mailbomber/internal/app/model"
)

type EmailScheduleRepository struct {
	store          *Store
	emailSchedules map[int]*model.EmailSchedule
}

func (r *EmailScheduleRepository) Create(es *model.EmailSchedule) error {
	if err := es.Validate(); err != nil {
		return err
	}

	es.ID = len(r.emailSchedules) + 1
	r.emailSchedules[es.ID] = es

	return nil
}

func (r *EmailScheduleRepository) SelectExecutables() ([]*model.EmailSchedule, error) {
	var schedules []*model.EmailSchedule
	for _, schedule := range r.emailSchedules {
		if schedule.ExecuteAfter.Before(time.Now()) {
			schedules = append(schedules, schedule)
		}
	}
	return schedules, nil
}

func (r *EmailScheduleRepository) UpdateExecutionTime(es *model.EmailSchedule) error {
	if schedule, ok := r.emailSchedules[es.ID]; ok {
		schedule.ExecuteAfter = time.Now().Add(time.Duration(es.ExecutionPeriod) * time.Hour)
		return nil
	}
	return fmt.Errorf("EmailSchedule with ID %d not found", es.ID)
}

func (r *EmailScheduleRepository) Delete(es *model.EmailSchedule) error {
	if _, ok := r.emailSchedules[es.ID]; ok {
		delete(r.emailSchedules, es.ID)
		return nil
	}
	return fmt.Errorf("EmailSchedule with ID %d not found", es.ID)
}
