package apiserver

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/nizepart/rest-go/internal/app"
	"github.com/nizepart/rest-go/internal/app/model"
	"gopkg.in/gomail.v2"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (s *server) handleEmailTemplateCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &model.EmailTemplate{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if err := s.store.EmailTemplate().Create(req); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.logger.Info("Email template created")
		s.respond(w, r, http.StatusCreated, req)
	}
}

func (s *server) handleEmailScheduleCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &model.EmailSchedule{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if err := s.store.EmailSchedule().Create(req); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.logger.Infof("Email schedule created to be executed after %v", req.ExecuteAfter)
		s.respond(w, r, http.StatusCreated, req)
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
	m.SetHeader("From", app.GetValue("SMTP_FROM", "noreply@localhost").String())
	m.SetHeader("To", req.To...)
	m.SetHeader("Cc", req.Cc...)
	m.SetHeader("Subject", req.Subject)
	m.SetBody(req.BodyType, req.Body)

	if req.Attach != "" {
		m.Attach(req.Attach)
	}

	s.logger.Infof("Sending email to %v", req.To)
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

		s.logger.Infof("Email sent to %v", req.To)
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
			s.logger.Debugf("Checking for email schedules. Found %d schedules", len(schedules))

			for _, schedule := range schedules {
				template, err := s.store.EmailTemplate().FindByID(schedule.EmailTemplateID)
				if err != nil {
					s.logger.Error(err)
					continue
				}

				recipients := strings.Split(schedule.Recipients, ",")

				msg := &model.Message{
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
