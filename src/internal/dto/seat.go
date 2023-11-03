package dto

import "github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/datastruct"

type UpdateSeatDTO struct {
	EventID uint              `json:"event_id,omitempty"`
	Status  datastruct.Status `json:"status,omitempty"`
}

type CreateSeatDTO struct {
	Status datastruct.Status `json:"status,omitempty"`
}
