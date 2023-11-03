package dto

type TicketAppBookingResponseDTO struct {
	BookingID  uint          `json:"booking_id"`
	CustomerID uint          `json:"customer_id"`
	Status     BookingStatus `json:"status"`
	Message    string        `json:"message"`
}

type ClientAppBookingRequestDTO struct {
	BookingID  uint `json:"id,omitempty"`
	CustomerID uint `json:"customer_id,omitempty"`
	EventID    uint `json:"event_id,omitempty"`
	SeatID     uint `json:"seat_id,omitempty"`
}

type BookingStatus string

const (
	BookingFailed    BookingStatus = "failed"
	BookingSuccess   BookingStatus = "success"
	BookingOnProcess BookingStatus = "on-going"
)
