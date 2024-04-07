package apiserver

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/nizepart/rest-go/model"
	"gopkg.in/gomail.v2"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (s *server) handleEmailTemplateCreate() http.HandlerFunc {
	type request struct {
		Subject  string `json:"subject"`
		Body     string `json:"body"`
		BodyType string `json:"body_type"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		et := &model.EmailTemplate{
			Subject:  req.Subject,
			Body:     req.Body,
			BodyType: req.BodyType,
		}
		if err := s.store.EmailTemplate().Create(et); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, et)
	}
}

func (s *server) handleEmailScheduleCreate() http.HandlerFunc {
	type request struct {
		EmailTemplateID int       `json:"email_template_id"`
		Recipients      string    `json:"recipients"`
		ExecuteAfter    time.Time `json:"execute_after"`
		ExecutionPeriod *int      `json:"execution_period"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		executionPeriod := 0
		if req.ExecutionPeriod != nil {
			executionPeriod = *req.ExecutionPeriod
		}

		es := &model.EmailSchedule{
			EmailTemplateID: req.EmailTemplateID,
			Recipients:      req.Recipients,
			ExecuteAfter:    req.ExecuteAfter,
			ExecutionPeriod: executionPeriod,
		}
		if err := s.store.EmailSchedule().Create(es); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, es)
	}
}

func (s *server) handleEmailTemplateGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			s.error(w, r, http.StatusBadRequest, errors.New("id not provided"))
			return
		}

		idInt, err := strconv.Atoi(id)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		et, err := s.store.EmailTemplate().FindByID(idInt)
		if err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}

		s.respond(w, r, http.StatusOK, et)
	}
}

func (s *server) sendEmail(req *model.Message) {
	m := gomail.NewMessage()
	m.SetHeader("From", req.From)
	m.SetHeader("To", req.To...)
	m.SetHeader("Cc", req.Cc...)
	m.SetHeader("Subject", req.Subject)
	m.SetBody(req.BodyType, req.Body)

	if req.Attach != "" {
		m.Attach(req.Attach)
	}

	s.emailService.Send(m)
}

func (s *server) handleEmailSend() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars["id"]

		req := &model.Message{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if ok {
			idInt, err := strconv.Atoi(id)
			if err != nil {
				s.error(w, r, http.StatusBadRequest, err)
				return
			}

			et, err := s.store.EmailTemplate().FindByID(idInt)
			if err != nil {
				s.error(w, r, http.StatusNotFound, err)
				return
			}

			req.Subject = et.Subject
			req.Body = et.Body
			req.BodyType = et.BodyType
		}

		s.sendEmail(req)

		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) startEmailScheduler() {
	go func() {
		for {
			schedules, err := s.store.EmailSchedule().SelectExecutables()
			if err != nil {
				s.logger.Error(err)
				continue
			}

			for _, schedule := range schedules {
				template, err := s.store.EmailTemplate().FindByID(schedule.EmailTemplateID)
				if err != nil {
					s.logger.Error(err)
					continue
				}

				recipients := strings.Split(schedule.Recipients, ",")

				msg := &model.Message{
					From:     "trapez@example.org",
					To:       recipients,
					Subject:  template.Subject,
					Body:     template.Body,
					BodyType: template.BodyType,
				}

				s.sendEmail(msg)

				if schedule.ExecutionPeriod > 0 {
					err = s.store.EmailSchedule().UpdateExecutionTime(schedule)
					if err != nil {
						s.logger.Error(err)
					}
				} else {
					err = s.store.EmailSchedule().Delete(schedule)
					if err != nil {
						s.logger.Error(err)
					}
				}
			}

			time.Sleep(1 * time.Minute)
		}
	}()
}
