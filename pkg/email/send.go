package email

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Send(to, subject, body string) error {
	b, _ := json.Marshal(struct {
		Personalizations []Personalization `json:"personalizations"`
		Content          []Content         `json:"content"`
		From             Contact           `json:"from"`
		ReplyTo          Contact           `json:"reply_to"`
	}{
		Personalizations: []Personalization{
			{
				To: []Contact{
					{
						Email: to,
						// TODO: Name
					},
				},
				Subject: subject,
			},
		},
		Content: []Content{
			{
				Type:  "text/plain",
				Value: body,
			},
		},
		From: Contact{
			Email: config.Sender,
		},
		ReplyTo: Contact{
			Email: config.Sender,
		},
	})
	url := "https://api.sendgrid.com/v3/mail/send"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return err
	}
	req.Header.Add("ContentType", "application/json")
	req.Header.Add("Authorization", "Bearer "+config.SendgridAPIKey)
	return nil
}

type Personalization struct {
	To      []Contact `json:"to"`
	Subject string    `json:"subject"`
}

type Contact struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Content struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
