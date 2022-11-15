package config

import (
	"crypto/tls"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

func SmtpClient() (*mail.SMTPClient, error) {
	// creating mail server
	var smtp *mail.SMTPServer = mail.NewSMTPClient()

	// configure mail server
	smtp.Host = "smtp.gmail.com"
	smtp.Username = ""
	smtp.Password = ""
	smtp.Port = 587
	smtp.SendTimeout = time.Second * 15
	smtp.ConnectTimeout = time.Second * 15
	smtp.Encryption = mail.EncryptionSTARTTLS

	// TODO: add TLS configuration or change TLSConfig
	// por el momento se esta obviando la verificacion de la encriptacion
	smtp.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// create mail client
	client, err := smtp.Connect()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	return client, nil
}
