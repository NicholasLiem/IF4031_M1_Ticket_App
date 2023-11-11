package emails

import (
	"fmt"
	"os"
)

func SendEmail(receiverEmailAddress string, subject string, htmlFilePath string) error {

	emailData := EmailData{
		Body: "This is the main content of the email.",
	}

	builder := NewEmailBuilder().
		From(os.Getenv("SENDER_EMAIL")).
		To([]string{receiverEmailAddress}).
		Subject(subject).
		HTMLBodyFromFile(htmlFilePath, emailData).
		AttachQRCode("booking-id")

	err := builder.Send()
	if err != nil {
		fmt.Printf("Error sending email: %s\n", err.Error())
		return err
	}
	return nil
}
