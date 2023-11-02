package datastruct

import "gorm.io/gorm"

type Seat struct {
	gorm.Model
	Status Status
}

type Status string

const (
	OPEN    Status = "open"
	ONGOING Status = "on-going"
	BOOKED  Status = "booked"
)
