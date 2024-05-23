package email_service

import (
	"log"
	"time"

	"github.com/nizepart/mailbomber/internal/app"
	"gopkg.in/gomail.v2"
)

type Service struct {
	ch chan *gomail.Message
}

func NewService() *Service {
	return &Service{
		ch: make(chan *gomail.Message),
	}
}

func (s *Service) Start() {
	go func() {
		d := gomail.NewDialer(app.GetValue("SMTP_HOST", "localhost").String(), app.GetValue("SMTP_PORT", "587").Int(), "", "")

		var sender gomail.SendCloser
		var err error
		open := false
		for {
			select {
			case m, ok := <-s.ch:
				if !ok {
					return
				}
				if !open {
					if sender, err = d.Dial(); err != nil {
						panic(err)
					}
					open = true
				}
				if err := gomail.Send(sender, m); err != nil {
					log.Print(err)
				}
			case <-time.After(30 * time.Second):
				if open {
					if err := sender.Close(); err != nil {
						panic(err)
					}
					open = false
				}
			}
		}
	}()
}

func (s *Service) Send(m *gomail.Message) {
	s.ch <- m
}

func (s *Service) Close() {
	close(s.ch)
}
