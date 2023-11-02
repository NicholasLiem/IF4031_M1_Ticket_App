package datastruct

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	CustomerID uint   `json:"customer_id,omitempty"`
	EventID    uint   `json:"event_id,omitempty"`
	SeatID     uint   `json:"seat_id,omitempty"`
	Status     string `json:"status,omitempty"`

	//Customer UserModel `gorm:"foreignKey:CustomerID"`
	Event Event `gorm:"foreignKey:EventID"`
	Seat  Seat  `gorm:"foreignKey:SeatID"`
}
