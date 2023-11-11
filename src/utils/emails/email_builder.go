package emails

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/utils"
	"github.com/jordan-wright/email"
	"html/template"
	"net/smtp"
	"os"
	"time"
)

// this fix of auth method, I get from here:
// https://gist.github.com/andelf/5118732

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("unknown from server")
		}
	}
	return nil, nil
}

type EmailBuilder struct {
	mail *email.Email
}

type EmailData struct {
	Timestamp string
	Body      string
}

func NewEmailBuilder() *EmailBuilder {
	return &EmailBuilder{
		mail: email.NewEmail(),
	}
}

func (eb *EmailBuilder) From(address string) *EmailBuilder {
	eb.mail.From = address
	return eb
}

func (eb *EmailBuilder) To(addresses []string) *EmailBuilder {
	eb.mail.To = addresses
	return eb
}

func (eb *EmailBuilder) Subject(subject string) *EmailBuilder {
	eb.mail.Subject = subject
	return eb
}

func (eb *EmailBuilder) HTMLBody(htmlContent string) *EmailBuilder {
	eb.mail.HTML = []byte(htmlContent)
	return eb
}

func (eb *EmailBuilder) HTMLBodyFromFile(templatePath string, data EmailData) *EmailBuilder {
	data.Timestamp = time.Now().Format("2006-01-02 15:04:05")

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Printf("Error parsing HTML template: %s\n", err.Error())
		return eb
	}

	var htmlContent bytes.Buffer
	err = tmpl.Execute(&htmlContent, data)
	if err != nil {
		fmt.Printf("Error executing HTML template: %s\n", err.Error())
		return eb
	}

	eb.mail.HTML = htmlContent.Bytes()
	return eb
}

func (eb *EmailBuilder) AttachQRCode(content string) *EmailBuilder {
	qrCodeBase64, err := utils.GenerateQRCode(content)
	if err != nil {
		fmt.Printf("Error generating QR Code: %s\n", err.Error())
		return eb
	}

	qrCodeData, err := base64.StdEncoding.DecodeString(qrCodeBase64)
	if err != nil {
		fmt.Printf("Error decoding QR Code: %s\n", err.Error())
		return eb
	}

	_, err = eb.mail.Attach(bytes.NewReader(qrCodeData), "qrCode.png", "image/png")
	if err != nil {
		fmt.Printf("Error attaching QR Code: %s\n", err.Error())
	}

	return eb
}

func (eb *EmailBuilder) Send() error {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	from := os.Getenv("SENDER_EMAIL")
	password := os.Getenv("SENDER_PASSWORD")
	auth := LoginAuth(from, password)
	return eb.mail.Send(smtpHost+":"+smtpPort, auth)
}
