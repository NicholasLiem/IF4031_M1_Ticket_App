package repository

import (
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/datastruct"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/dto"
	"gorm.io/gorm"
)

type EventQuery interface {
	CreateEvent(event datastruct.Event) (*datastruct.Event, error)
	UpdateEvent(eventID uint, event dto.UpdateEventDTO) (*datastruct.Event, error)
	DeleteEvent(eventID uint) (*datastruct.Event, error)
	GetEvent(eventID uint) (*datastruct.Event, error)
}

type eventQuery struct {
	pgdb *gorm.DB
}

func NewEventQuery(pgdb *gorm.DB) EventQuery {
	return &eventQuery{
		pgdb: pgdb,
	}
}

func (eq *eventQuery) CreateEvent(event datastruct.Event) (*datastruct.Event, error) {
	result := eq.pgdb.Create(&event)
	if result.Error != nil {
		return nil, result.Error
	}
	return &event, nil
}

func (eq *eventQuery) UpdateEvent(eventID uint, event dto.UpdateEventDTO) (*datastruct.Event, error) {
	existingEvent := datastruct.Event{}
	result := eq.pgdb.Where("id = ?", eventID).First(&existingEvent)
	if result.Error != nil {
		return nil, result.Error
	}

	existingEvent.EventName = event.EventName

	result = eq.pgdb.Save(&existingEvent)
	if result.Error != nil {
		return nil, result.Error
	}

	return &existingEvent, nil
}

func (eq *eventQuery) DeleteEvent(eventID uint) (*datastruct.Event, error) {
	existingEvent := datastruct.Event{}
	result := eq.pgdb.Where("id = ?", eventID).First(&existingEvent)
	if result.Error != nil {
		return nil, result.Error
	}

	result = eq.pgdb.Delete(&existingEvent)
	if result.Error != nil {
		return nil, result.Error
	}

	return &existingEvent, nil
}

func (eq *eventQuery) GetEvent(eventID uint) (*datastruct.Event, error) {
	event := datastruct.Event{}
	result := eq.pgdb.Where("id = ?", eventID).Preload("AvailableSeats").First(&event)
	if result.Error != nil {
		return nil, result.Error
	}

	return &event, nil
}
