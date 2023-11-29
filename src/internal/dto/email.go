package dto

type EmailMetaData struct {
	ReceiverEmailAddress string
	ReceiverName         string
	EmailSubject         string
	BodyMessage          string
	HTMLFilePath         string
	ContentDetails       EventPDF
}

func NewEmailMetaData(receiverEmailAddress, emailSubject, bodyMessage, htmlFilePath, fileName, eventName, seatID, userEmail, eventDate, qrContent string) *EmailMetaData {
	return &EmailMetaData{
		ReceiverEmailAddress: receiverEmailAddress,
		EmailSubject:         emailSubject,
		BodyMessage:          bodyMessage,
		HTMLFilePath:         htmlFilePath,
		ContentDetails: EventPDF{
			FileName:  fileName,
			EventName: eventName,
			SeatID:    seatID,
			UserEmail: userEmail,
			EventDate: eventDate,
			QRContent: qrContent,
		},
	}
}
