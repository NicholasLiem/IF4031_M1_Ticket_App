package repository

import (
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/datastruct"
	"gorm.io/gorm"
)

type SeatQuery interface {
	CreateSeat(seat datastruct.Seat) (*datastruct.Seat, error)
	UpdateSeat(seatID uint, seat datastruct.Seat) (*datastruct.Seat, error)
	DeleteSeat(seatID uint) (*datastruct.Seat, error)
	GetSeat(seatID uint) (*datastruct.Seat, error)
	GetSeatsForEvent(eventID uint) ([]datastruct.Seat, error)
}

type seatQuery struct {
	pgdb *gorm.DB
}

func NewSeatQuery(pgdb *gorm.DB) SeatQuery {
	return &seatQuery{
		pgdb: pgdb,
	}
}

func (sq *seatQuery) CreateSeat(seat datastruct.Seat) (*datastruct.Seat, error) {
	result := sq.pgdb.Create(&seat)
	if result.Error != nil {
		return nil, result.Error
	}
	
	return &seat, nil
}

func (sq *seatQuery) UpdateSeat(seatID uint, seat datastruct.Seat) (*datastruct.Seat, error) {
	existingSeat := datastruct.Seat{}
	result := sq.pgdb.First(&existingSeat, seatID)
	if result.Error != nil {
		return nil, result.Error
	}

	existingSeat.Status = seat.Status

	result = sq.pgdb.Save(&existingSeat)
	if result.Error != nil {
		return nil, result.Error
	}

	return &existingSeat, nil
}

func (sq *seatQuery) DeleteSeat(seatID uint) (*datastruct.Seat, error) {
	existingSeat := datastruct.Seat{}
	result := sq.pgdb.First(&existingSeat, seatID)
	if result.Error != nil {
		return nil, result.Error
	}

	result = sq.pgdb.Delete(&existingSeat)
	if result.Error != nil {
		return nil, result.Error
	}

	return &existingSeat, nil
}

func (sq *seatQuery) GetSeat(seatID uint) (*datastruct.Seat, error) {
	seat := datastruct.Seat{}
	result := sq.pgdb.First(&seat, seatID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &seat, nil
}

func (sq *seatQuery) GetSeatsForEvent(eventID uint) ([]datastruct.Seat, error) {
	var seats []datastruct.Seat
	result := sq.pgdb.Where("event_id = ?", eventID).Find(&seats)
	if result.Error != nil {
		return nil, result.Error
	}

	return seats, nil
}
