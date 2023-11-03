package service

import (
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/datastruct"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/repository"
)

type SeatService interface {
	CreateSeat(seat datastruct.Seat, eventID uint) (*datastruct.Seat, error)
	UpdateSeat(seatID uint, seat datastruct.Seat) (*datastruct.Seat, error)
	DeleteSeat(seatID uint) (*datastruct.Seat, error)
	GetSeat(seatID uint) (*datastruct.Seat, error)
	GetSeatsForEvent(eventID uint) ([]datastruct.Seat, error)
}

type seatService struct {
	dao repository.DAO
}

func NewSeatService(dao repository.DAO) SeatService {
	return &seatService{dao: dao}
}

func (ss *seatService) CreateSeat(seat datastruct.Seat, eventID uint) (*datastruct.Seat, error) {
	event, err := ss.dao.NewEventQuery().GetEvent(eventID)
	if err != nil && event != nil {
		return nil, err
	}

	seat.EventID = eventID
	return ss.dao.NewSeatQuery().CreateSeat(seat)
}

func (ss *seatService) UpdateSeat(seatID uint, seat datastruct.Seat) (*datastruct.Seat, error) {
	event, err := ss.dao.NewEventQuery().GetEvent(seat.EventID)
	if err != nil && event != nil {
		return nil, err
	}
	return ss.dao.NewSeatQuery().UpdateSeat(seatID, seat)
}

func (ss *seatService) DeleteSeat(seatID uint) (*datastruct.Seat, error) {
	return ss.dao.NewSeatQuery().DeleteSeat(seatID)
}

func (ss *seatService) GetSeat(seatID uint) (*datastruct.Seat, error) {
	return ss.dao.NewSeatQuery().GetSeat(seatID)
}

func (ss *seatService) GetSeatsForEvent(eventID uint) ([]datastruct.Seat, error) {
	event, err := ss.dao.NewEventQuery().GetEvent(eventID)
	if err != nil && event != nil {
		return nil, err
	}
	return ss.dao.NewSeatQuery().GetSeatsForEvent(eventID)
}
