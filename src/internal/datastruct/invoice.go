package datastruct

import uuid "github.com/satori/go.uuid"

type InvoiceRequest struct {
	BookingID  uuid.UUID `json:"booking_id,omitempty"`
	EventID    uint      `json:"event_id,omitempty"`
	CustomerID uint      `json:"customer_id,omitempty"`
	SeatID     uint      `json:"seat_id,omitempty"`
	Email      string    `json:"email,omitempty"`
}

type CancelInvoiceRequest struct{
	BookingID	uuid.UUID	`json:"booking_id,omitempty"`
}

type PaymentStatus string

const (
	SUCCESS PaymentStatus = "SUCCESS"
	PENDING PaymentStatus = "PENDING"
	FAILED  PaymentStatus = "FAILED"
)
