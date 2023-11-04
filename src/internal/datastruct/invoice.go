package datastruct

import "gorm.io/gorm"

type InvoiceRequest struct {
	BookingID  uint `json:"booking_id,omitempty"`
	EventID    uint `json:"event_id,omitempty"`
	CustomerID uint `json:"customer_id,omitempty"`
}

type Invoice struct {
	gorm.Model
	InvoiceID     string        `json:"id,omitempty"`
	BookingID     uint          `json:"booking_id,omitempty"`
	EventID       uint          `json:"event_id,omitempty"`
	CustomerID    uint          `json:"customer_id,omitempty"`
	PaymentURL    string        `json:"payment_url,omitempty"`
	PaymentStatus PaymentStatus `json:"payment_status,omitempty"`
}

type PaymentStatus string

const (
	PAID   PaymentStatus = "paid"
	UNPAID PaymentStatus = "unpaid"
)
