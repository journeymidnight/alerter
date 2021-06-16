package alerter

import (
	"gopkg.in/gomail.v2"
)

type Email struct {
	Dialer   *gomail.Dialer
	UserFrom string
	UserTo   []string
}

func InitEmailAlert(host string, port int, user string, password string, userFrom string, userTo []string) *Email {
	return &Email{
		Dialer:   gomail.NewDialer(host, port, user, password),
		UserFrom: userFrom,
		UserTo:   userTo,
	}
}

func (e *Email) Name() string {
	return "uos-email"
}

func (e *Email) Send(m Message) error {
	message := gomail.NewMessage()
	message.SetHeader("From", e.UserFrom)
	message.SetHeader("To", e.UserTo...)
	message.SetHeader("Subject", m.Type)
	message.SetBody("text/html", m.Info)
	sendCloser, err := e.Dialer.Dial()
	if err != nil {
		return err
	}
	defer sendCloser.Close()
	if err := gomail.Send(sendCloser, message); err != nil {
		return err
	}
	return nil
}

type EmailConfig struct {
	Enable   bool
	Host     string
	Port     int
	User     string
	Password string
	UserFrom string
	UserTo   []string
}

func (e EmailConfig) Type() AlertType {
	return EmailType
}