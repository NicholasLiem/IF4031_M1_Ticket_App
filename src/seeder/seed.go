package seeder

import (
	"log"
	"strings"

	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/datastruct"
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

const countEvents = 10
const seatPerEvent = 3

func Seed(db *gorm.DB) {
	// Seed Events
	SeedEvents(db, countEvents)

	// Seed Seats
	seedSeats(db, seatPerEvent)
}

func SeedEvents(db *gorm.DB, count int) {
	// Generate and seed fake event name
	events := make([]datastruct.Event, 0, count)

	// Create a set to store used event names
	usedNames := make(map[string]bool)

	for len(events) < count {
		// Generate a unique event name
		name := strings.Title("Event-" + gofakeit.State())
		if !usedNames[name] {
			usedNames[name] = true
			event := datastruct.Event{
				EventName: name,
			}
			events = append(events, event)
		}
	}

	if err := db.Create(&events).Error; err != nil {
		log.Fatalf("Failed to seed user: %v", err)
	}
}

func seedSeats(db *gorm.DB, count int) {
	var events int64

	// Fetch count of events from the database
	if err := db.Model(&datastruct.Event{}).Count(&events).Error; err != nil {
		log.Fatalf("Failed to fetch count of events: %v", err)
	}

	seats := make([]datastruct.Seat, 0, count*int(events))

	for i := 1; i <= int(events); i++ {
		for j := 1; j <= count; j++ {
			seat := datastruct.Seat{
				EventID: uint(i),
				Status:  datastruct.OPEN,
			}

			seats = append(seats, seat)
		}
	}

	if err := db.Create(&seats).Error; err != nil {
		log.Fatalf("Failed to seed seats: %v", err)
	}
}
