package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	EventName string    `gorm:"column:event_name;unique" json:"event_name,omitempty"`
	EventDate time.Time `gorm:"column:event_date" json:"event_date,omitempty"`
	Seats     []Seat    `json:"-"`
}
