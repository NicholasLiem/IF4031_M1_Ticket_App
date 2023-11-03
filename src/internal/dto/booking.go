package dto

type TicketAppBookingResponseDTO struct {
	BookingID uint          `json:"booking_id"`
	Status    BookingStatus `json:"status"`
	Message   string        `json:"message"`
}

type BookingStatus string

const (
	BookingFailed    BookingStatus = "failed"
	BookingSuccess   BookingStatus = "success"
	BookingOnProcess BookingStatus = "on-going"
)
