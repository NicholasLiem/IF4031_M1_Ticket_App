package datastruct

import "gorm.io/gorm"

// TODO: Make the routes
type Payment struct {
	gorm.Model
	InvoiceID          uint          `json:"invoice_id,omitempty"`
	Amount             float64       `json:"amount,omitempty"`
	PaymentStatus      PaymentStatus `json:"payment_status,omitempty"`
	TransactionDetails string        `json:"transaction_details,omitempty"`
	Invoice            Invoice       `gorm:"foreignKey:InvoiceID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type PaymentStatus string

const (
	PAID   PaymentStatus = "paid"
	UNPAID PaymentStatus = "unpaid"
)
