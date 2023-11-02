package datastruct

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	EventName      string
	AvailableSeats []Seat `gorm:"many2many:event_seats;"`
}
