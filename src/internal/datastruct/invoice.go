package datastruct

type InvoiceRequest struct {
	BookingID  uint `json:"booking_id,omitempty"`
	EventID    uint `json:"event_id,omitempty"`
	CustomerID uint `json:"customer_id,omitempty"`
	SeatID     uint `json:"seat_id,omitempty"`
}

type PaymentStatus string

const (
	SUCCESS PaymentStatus = "SUCCESS"
	FAILED  PaymentStatus = "FAILED"
)
