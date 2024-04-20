package sqlstore

import (
	"strconv"
	"time"

	"github.com/nizepart/rest-go/internal/app/model"
)

type EmailScheduleRepository struct {
	store *Store
}

func (r *EmailScheduleRepository) Create(es *model.EmailSchedule) error {
	if err := es.Validate(); err != nil {
		return err
	}

	return r.store.db.QueryRow("INSERT INTO email_schedule (email_template_id, recipients, execute_after, execution_period) VALUES ($1, $2, $3, $4) RETURNING id", es.EmailTemplateID, es.Recipients, es.ExecuteAfter, es.ExecutionPeriod).Scan(&es.ID)
}

func (r *EmailScheduleRepository) SelectExecutables() ([]*model.EmailSchedule, error) {
	// Query the database for email schedules where execute_after < time.Now()
	query := "SELECT * FROM email_schedule WHERE execute_after < $1"
	rows, err := r.store.db.Query(query, time.Now())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []*model.EmailSchedule
	for rows.Next() {
		s := &model.EmailSchedule{}
		if err := rows.Scan(&s.ID, &s.EmailTemplateID, &s.Recipients, &s.ExecuteAfter, &s.ExecutionPeriod); err != nil {
			return nil, err
		}
		schedules = append(schedules, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return schedules, nil
}

func (r *EmailScheduleRepository) UpdateExecutionTime(es *model.EmailSchedule) error {
	query := `UPDATE email_schedule SET execute_after = execute_after + ($1 || ' HOUR')::INTERVAL WHERE id = $2`
	_, err := r.store.db.Exec(query, strconv.Itoa(es.ExecutionPeriod), &es.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *EmailScheduleRepository) Delete(es *model.EmailSchedule) error {
	return r.store.db.QueryRow("DELETE FROM email_schedule WHERE id = $1", es.ID).Scan(&es.ID)
}
