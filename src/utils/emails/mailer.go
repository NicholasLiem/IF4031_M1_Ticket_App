package emails

import (
	"fmt"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/dto"
	"os"
)

func SendEmail(emailMetaData *dto.EmailMetaData, attachPDF bool) error {

	emailData := EmailData{
		Name: "Customer",
		Body: emailMetaData.BodyMessage,
	}

	builder := NewEmailBuilder().
		From(os.Getenv("SENDER_EMAIL")).
		To([]string{emailMetaData.ReceiverEmailAddress}).
		Subject(emailMetaData.EmailSubject).
		HTMLBodyFromFile(emailMetaData.HTMLFilePath, emailData)

	if attachPDF {
		builder = builder.AttachPDFWithQRCode(emailMetaData.ContentDetails, emailMetaData.ContentDetails.FileName)
	}

	err := builder.Send()
	if err != nil {
		fmt.Printf("Error sending email: %s\n", err.Error())
		return err
	}
	return nil
}
