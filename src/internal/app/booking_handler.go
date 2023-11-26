package app

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"

	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/datastruct"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/dto"
	response "github.com/NicholasLiem/IF4031_M1_Ticket_App/utils/http"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/utils/messages"
	uuid "github.com/satori/go.uuid"
)

func (m *MicroserviceServer) BookSeat(w http.ResponseWriter, r *http.Request) {
	// Prepare for response
	var responseDTO dto.BookingResponseDTO

	// Parse incoming request
	var requestDTO dto.IncomingBookingRequestDTO
	err := json.NewDecoder(r.Body).Decode(&requestDTO)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
		return
	}

	// Simulate a call that takes 2 seconds
	// time.Sleep(2 * time.Second)

	// Check the if event exists
	_, httpError := m.eventService.GetEvent(requestDTO.EventID)
	if httpError != nil {
		// sendBookingResponse(w, requestDTO.BookingID, requestDTO.CustomerID, requestDTO.EventID, requestDTO.SeatID, dto.BookingFailed, "Event tidak terdaftar", "", "", requestDTO.Email, responseDTO)
		response.ErrorResponse(w, httpError.StatusCode, httpError.Message)
		return
	}

	// Check the status of the seat
	existingSeat, err := m.seatService.GetSeat(requestDTO.SeatID)
	if err != nil || existingSeat.EventID != requestDTO.EventID {
		// sendBookingResponse(w, requestDTO.BookingID, requestDTO.CustomerID, requestDTO.EventID, requestDTO.SeatID, dto.BookingFailed, "Booking tidak dapat dilakukan. Kursi tidak terdaftar.", "", "", requestDTO.Email, responseDTO)
		response.ErrorResponse(w, http.StatusNotFound, "Booking can't be done. Seat does not exists.")
		return
	}

	//Return if the seat status does not open
	if existingSeat.Status != datastruct.OPEN {
		// sendBookingResponse(w, requestDTO.BookingID, requestDTO.CustomerID, requestDTO.EventID, requestDTO.SeatID, dto.BookingFailed, "Booking tidak dapat dilakukan. Kursi tidak open.", "", "", requestDTO.Email, responseDTO)
		response.ErrorResponse(w, http.StatusConflict, "Booking can't be done. Seat does not open.")
		return
	}

	// Simulate a 20% chance of failure
	if rand.Float32() < 0.2 {
		sendBookingResponse(w, requestDTO.BookingID, requestDTO.CustomerID, requestDTO.EventID, requestDTO.SeatID, dto.BookingFailed, "[Simulated Failure] Booking tidak dapat dilakukan.", uuid.Nil, "", requestDTO.Email, responseDTO)
		return
	}

	//Create Invoice To Payment App
	invoiceRequest := datastruct.InvoiceRequest{
		BookingID:  requestDTO.BookingID,
		CustomerID: requestDTO.CustomerID,
		EventID:    requestDTO.EventID,
		SeatID:     requestDTO.SeatID,
		Email:      requestDTO.Email,
	}

	requestBody, err := json.Marshal(invoiceRequest)
	if err != nil {
		// response.ErrorResponse(w, http.StatusInternalServerError, "[501] Error making invoice request")
		sendBookingResponse(w, requestDTO.BookingID, requestDTO.CustomerID, requestDTO.EventID, requestDTO.SeatID, dto.BookingFailed, "[501] Error making invoice request", uuid.Nil, "", requestDTO.Email, responseDTO)
		return
	}

	externalAPIPath := "/invoice"
	paymentResponse, err := m.restClientToPaymentApp.Post(externalAPIPath, requestBody)
	if err != nil {
		// response.ErrorResponse(w, http.StatusInternalServerError, "[502] Payment App is down")
		sendBookingResponse(w, requestDTO.BookingID, requestDTO.CustomerID, requestDTO.EventID, requestDTO.SeatID, dto.BookingFailed, "[502] Payment App is down", uuid.Nil, "", requestDTO.Email, responseDTO)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(paymentResponse.Body)

	if paymentResponse.StatusCode != http.StatusCreated {
		// response.ErrorResponse(w, http.StatusInternalServerError, "[503] Error making invoice request")
		sendBookingResponse(w, requestDTO.BookingID, requestDTO.CustomerID, requestDTO.EventID, requestDTO.SeatID, dto.BookingFailed, "[503] Error making invoice request", uuid.Nil, "", requestDTO.Email, responseDTO)
		return
	}

	dataBytes, err := response.GetJSONDataBytesFromResponse(paymentResponse)
	if err != nil {
		// response.ErrorResponse(w, http.StatusInternalServerError, "[504] Error making invoice request")
		sendBookingResponse(w, requestDTO.BookingID, requestDTO.CustomerID, requestDTO.EventID, requestDTO.SeatID, dto.BookingFailed, "[504] Error making invoice request", uuid.Nil, "", requestDTO.Email, responseDTO)
		return
	}

	var paymentResponseBody dto.IncomingPaymentResponseDTO
	if err := json.Unmarshal(dataBytes, &paymentResponseBody); err != nil {
		// response.ErrorResponse(w, http.StatusInternalServerError, "[505] Error making invoice request")
		sendBookingResponse(w, requestDTO.BookingID, requestDTO.CustomerID, requestDTO.EventID, requestDTO.SeatID, dto.BookingFailed, "[505] Error making invoice request", uuid.Nil, "", requestDTO.Email, responseDTO)
		return
	}

	//Set the seat status on hold and booking id
	existingSeat.Status = datastruct.ONGOING
	seat, err := m.seatService.UpdateSeat(existingSeat.ID, *existingSeat)
	if err != nil {
		// response.ErrorResponse(w, http.StatusInternalServerError, "Error updating seat status")
		sendBookingResponse(w, requestDTO.BookingID, requestDTO.CustomerID, requestDTO.EventID, requestDTO.SeatID, dto.BookingFailed, "Error updating seat status", uuid.Nil, "", requestDTO.Email, responseDTO)
		return
	}

	if seat.Status != datastruct.ONGOING {
		// response.ErrorResponse(w, http.StatusInternalServerError, "Error updating seat status")
		sendBookingResponse(w, requestDTO.BookingID, requestDTO.CustomerID, requestDTO.EventID, requestDTO.SeatID, dto.BookingFailed, "Error updating seat status", uuid.Nil, "", requestDTO.Email, responseDTO)
		return
	}

	// response.ErrorResponse(w, http.StatusConflict, "test.")
	// response.SuccessResponse(w, http.StatusOK, messages.SuccessfulDataObtain, nil)
	// return

	//Return Status on progress of the booking
	sendBookingResponse(w, requestDTO.BookingID, requestDTO.CustomerID, requestDTO.EventID, requestDTO.SeatID, dto.BookingOnProcess, "Booking success. Seat's status now is on-going.", paymentResponseBody.InvoiceID, paymentResponseBody.PaymentURL, requestDTO.Email, responseDTO)
	return
}

func (m *MicroserviceServer) CancelSeat(w http.ResponseWriter, r *http.Request) {
	// Prepare for response
	var requestDTO dto.IncomingCancelRequestDTO
	
	err := json.NewDecoder(r.Body).Decode(&requestDTO)
	if(err != nil){
		if err != nil {
			response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
			return
		}
	}
	invoiceRequest := datastruct.CancelInvoiceRequest{
		BookingID:  requestDTO.BookingID,
	}
	requestBody, err := json.Marshal(invoiceRequest)
	
	if err != nil {
		response.ErrorResponse(w,http.StatusBadRequest,"Invalid request body")
		return
	}
	externalAPIPath := "/invoice/cancel/" + invoiceRequest.BookingID.String()
	paymentResponse,err := m.restClientToPaymentApp.Put(externalAPIPath,requestBody)
	
	if(err != nil){
		response.ErrorResponse(w,http.StatusInternalServerError,"[500] Connection refused")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(paymentResponse.Body)
	// fmt.Println(paymentResponse)
	if paymentResponse.StatusCode != http.StatusOK {
		response.ErrorResponse(w,http.StatusNotFound,"[500] Error canceling invoice");
		return
	}

	dataBytes, err := response.GetJSONDataBytesFromResponse(paymentResponse)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, "[504] Error canceling invoice")
		return
	}

	var paymentResponseBody dto.IncomingPaymentResponseDTO
	if err := json.Unmarshal(dataBytes, &paymentResponseBody); err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, "[505] Error canceling invoice")
		return
	}
	response.SuccessResponse(w,http.StatusOK,"Your payment has been cancelled!",requestDTO.BookingID)
}
func sendBookingResponse(w http.ResponseWriter, bookingID uuid.UUID, customerID uint, eventID uint, seatID uint, status dto.BookingStatus, message string, invoiceID uuid.UUID, paymentURL string, email string, responseDTO dto.BookingResponseDTO) {
	responseDTO.BookingID = bookingID
	responseDTO.CustomerID = customerID
	responseDTO.EventID = eventID
	responseDTO.SeatID = seatID
	responseDTO.Status = status
	responseDTO.Message = message
	responseDTO.InvoiceID = invoiceID
	responseDTO.PaymentURL = paymentURL
	responseDTO.Email = email
	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulDataObtain, responseDTO)
}
