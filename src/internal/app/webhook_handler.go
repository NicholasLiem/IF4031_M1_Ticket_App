package app

import (
	"encoding/json"
	"fmt"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/datastruct"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/dto"
	response "github.com/NicholasLiem/IF4031_M1_Ticket_App/utils/http"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/utils/messages"
	"net/http"
)

func (m *MicroserviceServer) WebhookPaymentHandler(w http.ResponseWriter, r *http.Request) {
	var incomingInvoicePayload dto.IncomingInvoicePayload
	err := json.NewDecoder(r.Body).Decode(&incomingInvoicePayload)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
		return
	}

	//Check the status
	status := incomingInvoicePayload.PaymentStatus == datastruct.SUCCESS
	if !status {
		//Create PDF with QR Code
		fmt.Println("Making the QR...")
		fmt.Println("Making the PDF...")
		response.SuccessResponse(w, http.StatusOK, "Payment FAILED", nil)
		return
	}

	//Update seat status to booked if status success
	newSeatStatus := datastruct.Seat{
		Status: datastruct.BOOKED,
	}

	_, err = m.seatService.UpdateSeat(incomingInvoicePayload.SeatID, newSeatStatus)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Making the QR...")
		fmt.Println("Making the PDF...")
		response.SuccessResponse(w, http.StatusOK, "Update seat failed", nil)
		return
	}

	//Create PDF with QR Code
	fmt.Println("Making the QR...")
	fmt.Println("Making the PDF...")

	//Send to user email?
	fmt.Println(incomingInvoicePayload.BookingID)

	response.SuccessResponse(w, http.StatusOK, "Test", incomingInvoicePayload)
	return
}
