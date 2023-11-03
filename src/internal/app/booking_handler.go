package app

import (
	"encoding/json"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/datastruct"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/dto"
	response "github.com/NicholasLiem/IF4031_M1_Ticket_App/utils/http"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/utils/messages"
	"math/rand"
	"net/http"
	"time"
)

func (m *MicroserviceServer) BookSeat(w http.ResponseWriter, r *http.Request) {
	// Prepare for response
	var responseDTO dto.TicketAppBookingResponseDTO

	// Parse incoming request
	var requestDTO dto.ClientAppBookingRequestDTO
	err := json.NewDecoder(r.Body).Decode(&requestDTO)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
		return
	}

	// Simulate a call that takes 2 seconds
	time.Sleep(2 * time.Second)

	//Check the if event exists
	existingEvent, err := m.eventService.GetEvent(requestDTO.EventID)
	if err != nil || existingEvent == nil {
		sendBookingResponse(w, responseDTO, requestDTO.BookingID, requestDTO.CustomerID, dto.BookingFailed, "Event tidak terdaftar")
		return
	}

	//Check the status of the seat
	existingSeat, err := m.seatService.GetSeat(requestDTO.SeatID)
	if err != nil {
		sendBookingResponse(w, responseDTO, requestDTO.BookingID, requestDTO.CustomerID, dto.BookingFailed, "Booking tidak dapat dilakukan. Kursi tidak terdaftar.")
		return
	}

	//Return if the seat status is not open
	if existingSeat.Status != datastruct.OPEN {
		sendBookingResponse(w, responseDTO, requestDTO.BookingID, requestDTO.CustomerID, dto.BookingFailed, "Booking tidak dapat dilakukan. Kursi tidak open.")
		return
	}

	// Simulate a 20% chance of failure
	if rand.Float32() < 0.2 {
		sendBookingResponse(w, responseDTO, requestDTO.BookingID, requestDTO.CustomerID, dto.BookingFailed, "[Simulated Failure] Booking tidak dapat dilakukan.")
		return
	}

	//Set the seat status on hold and booking id
	existingSeat.Status = datastruct.ONGOING
	seat, err := m.seatService.UpdateSeat(existingSeat.ID, *existingSeat)
	if err != nil {
		return
	}

	if seat.Status != datastruct.ONGOING {
		response.ErrorResponse(w, http.StatusInternalServerError, "Error updating seat status")
		return
	}

	//Return Status on progress of the booking
	sendBookingResponse(w, responseDTO, requestDTO.BookingID, requestDTO.CustomerID, dto.BookingOnProcess, "Booking kursi berhasil dilakukan. Status kursi sekarang on-going.")
	return
}

func sendBookingResponse(w http.ResponseWriter, responseDTO dto.TicketAppBookingResponseDTO, bookingId uint, customerId uint, status dto.BookingStatus, message string) {
	responseDTO.BookingID = bookingId
	responseDTO.CustomerID = customerId
	responseDTO.Status = status
	responseDTO.Message = message
	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulDataObtain, responseDTO)
}
