package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/datastruct"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/dto"
	"github.com/redis/go-redis/v9"
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
	pgdb  *gorm.DB
	redis *redis.Client
}

func NewEventQuery(pgdb *gorm.DB, redis *redis.Client) EventQuery {
	return &eventQuery{
		pgdb:  pgdb,
		redis: redis,
	}
}

func (eq *eventQuery) CreateEvent(event datastruct.Event) (*datastruct.Event, error) {
	result := eq.pgdb.Create(&event)
	if result.Error != nil {
		return nil, result.Error
	}

	ctx := context.Background()
	cacheKey := fmt.Sprintf("event:%d", event.ID)
	marshaledEvent, err := json.Marshal(event)
	if err == nil {
		_ = eq.redis.Set(ctx, cacheKey, marshaledEvent, time.Hour).Err()
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

	ctx := context.Background()
	cacheKey := fmt.Sprintf("event:%d", eventID)
	_ = eq.redis.Del(ctx, cacheKey).Err()

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

	ctx := context.Background()
	cacheKey := fmt.Sprintf("event:%d", eventID)
	_ = eq.redis.Del(ctx, cacheKey).Err()

	return &existingEvent, nil
}

func (eq *eventQuery) GetEvent(eventID uint) (*datastruct.Event, error) {
	ctx := context.Background()
	event := datastruct.Event{}

	cacheKey := fmt.Sprintf("event:%d", event.ID)
	_ = eq.redis.Del(ctx, cacheKey).Err()

	result := eq.pgdb.Where("id = ?", eventID).Preload("Seats").First(&event)
	if result.Error != nil {
		return nil, result.Error
	}

	marshaledEvent, err := json.Marshal(event)
	if err == nil {
		// Set the cache with an expiration time (e.g., 1 hour)
		_ = eq.redis.Set(ctx, cacheKey, marshaledEvent, time.Hour).Err()
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
