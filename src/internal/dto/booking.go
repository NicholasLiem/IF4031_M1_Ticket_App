package dto

type BookingResponseDTO struct {
	InvoiceID  string        `json:"invoice_id,omitempty"`
	BookingID  uint          `json:"booking_id,omitempty"`
	EventID    uint          `json:"event_id,omitempty"`
	SeatID     uint          `json:"seat_id,omitempty"`
	Email      string        `json:"email,omitempty"`
	CustomerID uint          `json:"customer_id,omitempty"`
	PaymentURL string        `json:"payment_uRL,omitempty"`
	Status     BookingStatus `json:"status,omitempty"`
	Message    string        `json:"message,omitempty"`
}

type IncomingBookingRequestDTO struct {
	BookingID  uint   `json:"booking_id,omitempty"`
	CustomerID uint   `json:"customer_id,omitempty"`
	EventID    uint   `json:"event_id,omitempty"`
	SeatID     uint   `json:"seat_id,omitempty"`
	Email      string `json:"email,omitempty"`
}

type IncomingPaymentResponseDTO struct {
	InvoiceID  string `json:"id,omitempty"`
	CustomerID uint   `json:"customerID,omitempty"`
	BookingID  uint   `json:"bookingID,omitempty"`
	EventID    uint   `json:"eventID,omitempty"`
	SeatID     uint   `json:"seatID,omitempty"`
	Email      string `json:"email,omitempty"`
	PaymentURL string `json:"paymentURL,omitempty"`
}

type BookingStatus string

const (
	BookingFailed    BookingStatus = "failed"
	BookingSuccess   BookingStatus = "success"
	BookingOnProcess BookingStatus = "on-going"
)
