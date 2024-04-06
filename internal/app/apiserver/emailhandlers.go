package apiserver

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/nizepart/rest-go/model"
	"gopkg.in/gomail.v2"
	"net/http"
	"strconv"
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

		s.respond(w, r, http.StatusOK, nil)
	}
}
