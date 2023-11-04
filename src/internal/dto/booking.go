package dto

//TODO: organize ini dto"nya

type TicketAppBookingResponseDTO struct {
	InvoiceID  string        `json:"invoice_id,omitempty"`
	BookingID  uint          `json:"booking_id,omitempty"`
	EventID    uint          `json:"event_id,omitempty"`
	SeatID     uint          `json:"seat_id,omitempty"`
	CustomerID uint          `json:"customer_id,omitempty"`
	PaymentURL string        `json:"payment_uRL,omitempty"`
	Status     BookingStatus `json:"status,omitempty"`
	Message    string        `json:"message,omitempty"`
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
	SeatID     uint   `json:"seatID,omitempty"`
	CustomerID uint   `json:"customerID,omitempty"`
	PaymentURL string `json:"paymentURL,omitempty"`
}

type BookingStatus string

const (
	BookingFailed    BookingStatus = "failed"
	BookingSuccess   BookingStatus = "success"
	BookingOnProcess BookingStatus = "on-going"
)
