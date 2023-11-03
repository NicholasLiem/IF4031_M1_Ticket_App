package datastruct

import "gorm.io/gorm"

// TODO: Make the routes
type Invoice struct {
	gorm.Model
	BookingID     uint          `json:"booking_id,omitempty"`
	TotalAmount   float64       `json:"total_amount,omitempty"`
	PaymentStatus PaymentStatus `json:"payment_status,omitempty"`
	DueDate       string        `json:"due_date,omitempty"`
	Booking       Booking       `gorm:"foreignKey:BookingID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
