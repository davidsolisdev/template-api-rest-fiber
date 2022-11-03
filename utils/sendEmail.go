package utils

import "net/smtp"

type Email struct {
	From     string
	Password string
	To       []string
	Message  string
	HostSMTP string
	PortSMTP string
}

func SendEmail(email *Email) (err error) {
	if len(email.HostSMTP) < 5 {
		email.HostSMTP = "smtp.gmail.com"
	}

	if len(email.PortSMTP) < 2 {
		email.PortSMTP = "587"
	}

	var auth smtp.Auth = smtp.PlainAuth("", email.From, email.Password, email.HostSMTP)

	err = smtp.SendMail(email.HostSMTP+":"+email.PortSMTP, auth, email.From, email.To, []byte(email.Message))

	return err
}
