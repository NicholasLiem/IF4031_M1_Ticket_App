package dto

import "github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/datastruct"

type IncomingInvoicePayload struct {
	InvoiceID     string                   `json:"id,omitempty"`
	BookingID     uint                     `json:"bookingID,omitempty"`
	EventID       uint                     `json:"eventID,omitempty"`
	SeatID        uint                     `json:"seatID,omitempty"`
	Email         string                   `json:"email,omitempty"`
	CustomerID    uint                     `json:"customerID,omitempty"`
	PaymentURL    string                   `json:"paymentURL,omitempty"`
	PaymentStatus datastruct.PaymentStatus `json:"paymentStatus,omitempty"`
}
