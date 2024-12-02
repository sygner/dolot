package pkg

import (
	"bytes"
	"dolott_authentication/internal/types"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/smtp"
)

const (
	EMAIL_USERNAME  = "alexjjx8@gmail.com"
	EMAIL_PASSWORD  = "nqogrfuyjelpvnjd"
	SMTP_SERVER     = "smtp.gmail.com"
	GMAIL_SMTP_PORT = "587"
)

type Email struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	From    string `json:"email"`
	Auth    smtp.Auth
}

func NewEmail(to, subject string) *Email {
	auth := smtp.PlainAuth("", EMAIL_USERNAME, EMAIL_PASSWORD, SMTP_SERVER)
	return &Email{
		To:      to,
		Subject: subject,
		From:    EMAIL_USERNAME,
		Auth:    auth,
	}
}

func (c *Email) SendAuthEmail(hbsFilePath, code string) *types.Error {
	htmlContent, err := ioutil.ReadFile(hbsFilePath)
	if err != nil {
		fmt.Println(err)
		return types.NewInternalError("internal issue, error code 9003")
	}

	tmpl, err := template.New("emailTemplate").Parse(string(htmlContent))
	if err != nil {
		return types.NewInternalError("internal issue, error code 9004")
	}

	data := struct {
		Name string
	}{
		Name: code,
	}

	var body bytes.Buffer
	err = tmpl.Execute(&body, data)
	if err != nil {
		return types.NewInternalError("internal issue, error code 9005")
	}

	msg := bytes.Buffer{}
	msg.WriteString(fmt.Sprintf("From: %s\r\n", c.From))
	msg.WriteString(fmt.Sprintf("To: %s\r\n", c.To))
	msg.WriteString(fmt.Sprintf("Subject: %s\r\n", c.Subject))
	msg.WriteString("MIME-Version: 1.0\r\n")
	msg.WriteString("Content-Type: text/html; charset=\"UTF-8\"\r\n")
	msg.WriteString("\r\n")
	msg.WriteString(body.String())

	err = smtp.SendMail(
		SMTP_SERVER+":"+GMAIL_SMTP_PORT,
		c.Auth,
		c.From,
		[]string{c.To},
		msg.Bytes(),
	)

	if err != nil {
		// Consider logging the error or passing it to an error handler
		fmt.Println("Error sending email:", err)
		return types.NewInternalError("internal issue, error code 9006")
	}
	fmt.Println("Email sent successfully!")
	return nil
}
