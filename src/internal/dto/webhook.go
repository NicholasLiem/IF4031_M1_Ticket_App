package dto

import (
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/datastruct"
	uuid "github.com/satori/go.uuid"
)

type IncomingInvoicePayload struct {
	InvoiceID     uuid.UUID                `json:"id,omitempty"`
	BookingID     uuid.UUID                `json:"bookingID,omitempty"`
	EventID       uint                     `json:"eventID,omitempty"`
	SeatID        uint                     `json:"seatID,omitempty"`
	Email         string                   `json:"email,omitempty"`
	CustomerID    uint                     `json:"customerID,omitempty"`
	PaymentURL    string                   `json:"paymentURL,omitempty"`
	PaymentStatus datastruct.PaymentStatus `json:"paymentStatus,omitempty"`
}

type InvoicePayloadRequestToClient struct {
	InvoiceID     uuid.UUID                `json:"id,omitempty"`
	BookingID     uuid.UUID                `json:"bookingID,omitempty"`
	PaymentURL    string                   `json:"paymentURL,omitempty"`
	PaymentStatus datastruct.PaymentStatus `json:"paymentStatus,omitempty"`
	Message       string                   `json:"message,omitempty"`
}
