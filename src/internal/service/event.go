package service

import (
	"net/http"

	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/datastruct"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/dto"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/repository"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/utils"
)

type EventService interface {
	CreateEvent(event datastruct.Event) (*datastruct.Event, *utils.HttpError)
	UpdateEvent(eventID uint, event dto.UpdateEventDTO) (*datastruct.Event, *utils.HttpError)
	DeleteEvent(eventID uint) (*datastruct.Event, *utils.HttpError)
	GetEvent(eventID uint) (*datastruct.Event, *utils.HttpError)
}

type eventService struct {
	dao repository.DAO
}

func NewEventService(dao repository.DAO) EventService {
	return &eventService{dao: dao}
}

func (es *eventService) CreateEvent(event datastruct.Event) (*datastruct.Event, *utils.HttpError) {
	// Check unique constraint
	isUnique, err := es.dao.NewEventQuery().IsEventNameUnique(event.EventName, nil)

	if err != nil {
		return nil, &utils.HttpError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	if !isUnique {
		return nil, &utils.HttpError{
			Message:    "Event name is already exists",
			StatusCode: http.StatusConflict,
		}
	}

	createdEvent, err := es.dao.NewEventQuery().CreateEvent(event)
	if err != nil {
		return nil, &utils.HttpError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return createdEvent, nil
}

func (es *eventService) UpdateEvent(eventID uint, event dto.UpdateEventDTO) (*datastruct.Event, *utils.HttpError) {
	// Check unique constraint
	isUnique, err := es.dao.NewEventQuery().IsEventNameUnique(event.EventName, &eventID)

	if err != nil {
		return nil, &utils.HttpError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	if !isUnique {
		return nil, &utils.HttpError{
			Message:    "Event name is already exists",
			StatusCode: http.StatusConflict,
		}
	}

	updatedEvent, err := es.dao.NewEventQuery().UpdateEvent(eventID, event)
	if err != nil {
		return nil, &utils.HttpError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return updatedEvent, nil
}

func (es *eventService) DeleteEvent(eventID uint) (*datastruct.Event, *utils.HttpError) {
	eventData, err := es.dao.NewEventQuery().GetEvent(eventID)
	if err != nil && eventData != nil {
		return nil, &utils.HttpError{
			Message:    "Event not found",
			StatusCode: http.StatusNotFound,
		}
	}
	deletedEvent, err := es.dao.NewEventQuery().DeleteEvent(eventID)
	if err != nil {
		return nil, &utils.HttpError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return deletedEvent, nil
}

func (es *eventService) GetEvent(eventID uint) (*datastruct.Event, *utils.HttpError) {
	eventData, err := es.dao.NewEventQuery().GetEvent(eventID)
	if err != nil && eventData != nil {
		return nil, &utils.HttpError{
			Message:    "Event not found",
			StatusCode: http.StatusNotFound,
		}
	}

	return eventData, nil
}
