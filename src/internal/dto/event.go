package dto

type UpdateEventDTO struct {
	EventName string `json:"event_name,omitempty"`
}

type CreateEventDTO struct {
	EventName string `json:"event_name,omitempty"`
}
