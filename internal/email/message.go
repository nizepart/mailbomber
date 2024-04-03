package email

import validation "github.com/go-ozzo/ozzo-validation"

type Message struct {
	From     string   `json:"from"`
	To       []string `json:"to"`
	Cc       []string `json:"cc"`
	Subject  string   `json:"subject"`
	Body     string   `json:"body"`
	BodyType string   `json:"bodyType"`
	Attach   string   `json:"attach"`
}

func (m *Message) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.From, validation.Required, validation.Length(1, 100)),
		validation.Field(&m.To, validation.Required, validation.Length(1, 100)),
		validation.Field(&m.Cc, validation.Length(0, 100)),
		validation.Field(&m.Subject, validation.Required, validation.Length(1, 100)),
		validation.Field(&m.Body, validation.Required),
		validation.Field(&m.BodyType, validation.Required, validation.In("text/plain", "text/html")),
		validation.Field(&m.Attach, validation.Length(0, 255)),
	)
}
