package dto

type EventPDF struct {
	FileName  string
	EventName string
	SeatID    string
	UserEmail string
	EventDate string
	QRContent string
}

func NewEventPDF(fileName, eventName, seatID, userEmail, eventDate, qrContent string) *EventPDF {
	return &EventPDF{
		FileName:  fileName,
		EventName: eventName,
		SeatID:    seatID,
		UserEmail: userEmail,
		EventDate: eventDate,
		QRContent: qrContent,
	}
}
