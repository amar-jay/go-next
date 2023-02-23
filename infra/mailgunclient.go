package mailgunclient

import (
	"github.com/amar-jay/go-api-boilerplate/utils/config"
	"github.com/mailgun/mailgun-go/v4"
)

type MailgunClient interface {
	Welcome(subject, text, to, htmlStr string) error
	ResetPassword(subject, text, to, htmlStr string) error
}

type mailgunClient struct {
	conf   config.Config
	client *mailgun.MailgunImpl
}

func NewMailgunClient(c config.Config) *mailgunClient {
	return &mailgunClient{
		conf:   c,
		client: mailgun.NewMailgun(c.Mailgun.Domain, c.Mailgun.APIKey),
	}
}
