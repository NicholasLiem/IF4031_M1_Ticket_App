package dto

type TicketAppBookingResponseDTO struct {
	BookingID  uint          `json:"booking_id"`
	CustomerID uint          `json:"customer_id"`
	InvoiceID  string        `json:"invoice_id"`
	PaymentURL string        `json:"payment_url"`
	Status     BookingStatus `json:"status"`
	Message    string        `json:"message"`
}

type ClientAppBookingRequestDTO struct {
	BookingID  uint `json:"booking_id,omitempty"`
	CustomerID uint `json:"customer_id,omitempty"`
	EventID    uint `json:"event_id,omitempty"`
	SeatID     uint `json:"seat_id,omitempty"`
}

type PaymentResponseDTO struct {
	InvoiceID  string `json:"id,omitempty"`
	BookingID  uint   `json:"bookingID,omitempty"`
	EventID    uint   `json:"eventID,omitempty"`
	CustomerID uint   `json:"customerID,omitempty"`
	PaymentURL string `json:"paymentURL,omitempty"`
}

type BookingStatus string

const (
	BookingFailed    BookingStatus = "failed"
	BookingSuccess   BookingStatus = "success"
	BookingOnProcess BookingStatus = "on-going"
)
