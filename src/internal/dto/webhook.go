package dto

import "github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/datastruct"

type IncomingInvoicePayload struct {
	InvoiceID     string                   `json:"invoice_id,omitempty"`
	BookingID     uint                     `json:"booking_id,omitempty"`
	EventID       uint                     `json:"event_id,omitempty"`
	SeatID        uint                     `json:"seat_id,omitempty"`
	CustomerID    uint                     `json:"customer_id,omitempty"`
	PaymentURL    string                   `json:"payment_url,omitempty"`
	PaymentStatus datastruct.PaymentStatus `json:"payment_status,omitempty"`
}
