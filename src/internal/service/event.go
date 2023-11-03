package service

import (
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/datastruct"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/dto"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/repository"
)

type EventService interface {
	CreateEvent(event datastruct.Event) (*datastruct.Event, error)
	UpdateEvent(eventID uint, event dto.UpdateEventDTO) (*datastruct.Event, error)
	DeleteEvent(eventID uint) (*datastruct.Event, error)
	GetEvent(eventID uint) (*datastruct.Event, error)
}

type eventService struct {
	dao repository.DAO
}

func NewEventService(dao repository.DAO) EventService {
	return &eventService{dao: dao}
}

func (es *eventService) CreateEvent(event datastruct.Event) (*datastruct.Event, error) {
	return es.dao.NewEventQuery().CreateEvent(event)
}

func (es *eventService) UpdateEvent(eventID uint, event dto.UpdateEventDTO) (*datastruct.Event, error) {
	return es.dao.NewEventQuery().UpdateEvent(eventID, event)
}

func (es *eventService) DeleteEvent(eventID uint) (*datastruct.Event, error) {
	return es.dao.NewEventQuery().DeleteEvent(eventID)
}

func (es *eventService) GetEvent(eventID uint) (*datastruct.Event, error) {
	return es.dao.NewEventQuery().GetEvent(eventID)
}
