package datastruct

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	EventName string `gorm:"column:event_name" json:"event_name,omitempty"`
	Seats     []Seat `json:"-"`
}
