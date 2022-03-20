package mailer

import (
	"bytes"
	"github.com/go-mail/mail"
	"github.com/rhodeon/prettylog"
	"html/template"
	"time"
)

type Mailer struct {
	dialer *mail.Dialer
	sender string
}

func New(host string, port int, username string, password string) *Mailer {
	const sender = "Team Sniphub <no-reply@sniphub.mail.com>"
	dialer := mail.NewDialer(host, port, username, password)
	dialer.Timeout = 0 * time.Second

	return &Mailer{
		dialer: dialer,
		sender: sender,
	}
}

func (m *Mailer) Send(recipient string, templateFile string, data interface{}) error {
	tmpl, err := template.New("email").ParseFiles(templateFile)
	if err != nil {
		return err
	}

	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return err
	}
	prettylog.InfoF("SUBJECT: %s", subject)

	plainBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(plainBody, "plainBody", data)
	if err != nil {
		return err
	}
	prettylog.InfoF("PLAIN: %s", plainBody)

	htmlBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(htmlBody, "htmlBody", data)
	if err != nil {
		return err
	}
	prettylog.InfoF("HTML: %s", htmlBody)

	msg := mail.NewMessage()
	msg.SetHeader("To", recipient)
	msg.SetHeader("From", m.sender)
	msg.SetHeader("Subject", subject.String())
	msg.SetBody("text/plain", plainBody.String())
	msg.AddAlternative("text/html", htmlBody.String())

	err = m.dialer.DialAndSend(msg)
	if err != nil {
		return err
	}

	return nil
}
