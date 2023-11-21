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
	IsEventNameUnique(eventName string, eventID *uint) (bool, error)
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
	result := eq.pgdb.Where("id = ?", eventID).Preload("Seats").First(&event)
	if result.Error != nil {
		return nil, result.Error
	}

	return &event, nil
}

func (u *eventQuery) IsEventNameUnique(eventName string, eventID *uint) (bool, error) {
	// If eventID is provided, exclude it from the query
	query := u.pgdb.Model(datastruct.Event{}).Where("event_name = ?", eventName)
	if eventID != nil {
		query = query.Where("id <> ?", *eventID)
	}

	// Execute the query to check event name uniqueness
	var count int64
	if err := query.Count(&count).Error; err != nil {
		return false, err
	}

	// If count is 0, event is unique; otherwise, it's not
	return count == 0, nil
}
