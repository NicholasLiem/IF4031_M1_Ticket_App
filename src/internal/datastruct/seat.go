package datastruct

import "gorm.io/gorm"

type Seat struct {
	gorm.Model
	EventID uint   `gorm:"column:event_id" json:"event_id,omitempty"`
	Status  Status `gorm:"column:status" json:"status,omitempty"`
	Event   Event  `gorm:"foreignKey:EventID" json:"-"`
}

type Status string

const (
	OPEN    Status = "open"
	ONGOING Status = "on-going"
	BOOKED  Status = "booked"
)

func IsValidStatus(status Status) bool {
	switch status {
	case OPEN, ONGOING, BOOKED:
		return true
	default:
		return false
	}
}
