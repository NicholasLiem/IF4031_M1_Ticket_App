package app

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"time"

	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/datastruct"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/dto"
	response "github.com/NicholasLiem/IF4031_M1_Ticket_App/utils/http"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/utils/messages"
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
		// response.SuccessResponse(w, http.StatusOK, "Payment FAILED", nil)

		// Call webhook in Client App
		// Create Invoice To Client
		invoiceRequest := dto.InvoicePayloadRequestToClient{
			InvoiceID:     incomingInvoicePayload.InvoiceID,
			BookingID:     incomingInvoicePayload.BookingID,
			PaymentURL:    incomingInvoicePayload.PaymentURL,
			PaymentStatus: incomingInvoicePayload.PaymentStatus,
			Message:       "Payment is failed, please kindly check your email",
		}

		requestBody, err := json.Marshal(invoiceRequest)
		if err != nil {
			response.ErrorResponse(w, http.StatusInternalServerError, "[501] Error making invoice request")
			return
		}

		externalAPIPath := "/webhook"
		paymentResponse, err := m.restClientToClientApp.Put(externalAPIPath, requestBody)
		if err != nil {
			response.ErrorResponse(w, http.StatusInternalServerError, "[502] Payment App is down")
			return
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				return
			}
		}(paymentResponse.Body)

		if paymentResponse.StatusCode != http.StatusOK {
			response.ErrorResponse(w, http.StatusInternalServerError, "Webhook response: "+paymentResponse.Status)
			return
		}

		response.SuccessResponse(w, http.StatusOK, "Webhook response: "+paymentResponse.Status, nil)
		return
	}

	//Update seat status to booked if status success
	newSeatStatus := datastruct.Seat{
		Status: datastruct.BOOKED,
	}

	// Retry attempt if update seat status is failed
	// Simple mechanism
	attemptCount := 3

	for i := 0; i < attemptCount; i++ {
		_, updateSeatErr := m.seatService.UpdateSeat(incomingInvoicePayload.SeatID, newSeatStatus)
		if updateSeatErr == nil {
			// If seat update is successful or no error occurs, break out of the loop
			break
		}

		fmt.Println(err)
		fmt.Println("Making the QR...")
		fmt.Println("Making the PDF...")
		fmt.Println("Update seat failed, retrying...")

		// Wait for a while before retrying
		time.Sleep(time.Duration(math.Pow(2, float64(i))) * time.Second)
	}

	// _, err = m.seatService.UpdateSeat(incomingInvoicePayload.SeatID, newSeatStatus)
	// if err != nil {
	// 	fmt.Println(err)
	// 	fmt.Println("Making the QR...")
	// 	fmt.Println("Making the PDF...")
	// 	response.SuccessResponse(w, http.StatusOK, "Update seat failed", nil)
	// }

	//Create PDF with QR Code
	fmt.Println("Making the QR...")
	fmt.Println("Making the PDF...")

	//Send to user email?
	fmt.Println(incomingInvoicePayload.BookingID)

	// response.SuccessResponse(w, http.StatusOK, "Test", incomingInvoicePayload)

	// Call webhook in Client App
	// Create Invoice To Client
	invoiceRequest := dto.InvoicePayloadRequestToClient{
		InvoiceID:     incomingInvoicePayload.InvoiceID,
		BookingID:     incomingInvoicePayload.BookingID,
		PaymentURL:    incomingInvoicePayload.PaymentURL,
		PaymentStatus: incomingInvoicePayload.PaymentStatus,
		Message:       "Payment is successful, please kindly check your email",
	}

	requestBody, err := json.Marshal(invoiceRequest)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, "[501] Error making invoice request")
		return
	}

	externalAPIPath := "/webhook"
	paymentResponse, err := m.restClientToClientApp.Put(externalAPIPath, requestBody)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, "[502] Payment App is down")
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(paymentResponse.Body)

	if paymentResponse.StatusCode != http.StatusOK {
		response.ErrorResponse(w, paymentResponse.StatusCode, "Webhook response: "+paymentResponse.Status)
		return
	}

	response.SuccessResponse(w, http.StatusOK, "Webhook response: "+paymentResponse.Status, nil)
	return
}
