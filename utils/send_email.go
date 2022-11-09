package utils

import (
	"github.com/davidsolisdev/template-api-rest-fiber/config"

	mail "github.com/xhit/go-simple-mail/v2"
)

type NewEmail struct {
	From    string
	To      string
	Subject string
}

func SendEmail(email *NewEmail, bodyHtml string) (bool, error) {
	// configure default data
	if len(email.From) < 5 {
		email.From = "example@gmail.com"
	}
	// conecting mail server
	smtp, err := config.SmtpClient()
	if err != nil {
		return false, err
	}

	// create new mail
	var eMail *mail.Email = mail.NewMSG()
	eMail.SetFrom(email.From)
	eMail.AddTo(email.To)
	eMail.SetSubject(email.Subject)
	eMail.SetBody(mail.TextHTML, bodyHtml)

	// comprobate mail
	if eMail.Error != nil {
		return false, eMail.Error
	}

	// send mail
	err = eMail.Send(smtp)
	if err != nil {
		return false, err
	}

	return true, err
}
