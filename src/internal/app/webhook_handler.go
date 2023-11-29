package app

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/NicholasLiem/IF4031_M1_Ticket_App/utils/emails"

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

	//find event name
	eventData, httpError := m.eventService.GetEvent(incomingInvoicePayload.EventID)
	if httpError != nil {
		response.ErrorResponse(w, http.StatusNotFound, "Event data not found for this event id")
		return
	}

	layout := "2006-01-02 15:04:05.999999 -0700 MST"
	parsedTime, err := time.Parse(layout, eventData.CreatedAt.String())
	if err != nil {
		fmt.Println("Error parsing time:", err)
		response.ErrorResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	formattedDate := parsedTime.Format("2006-01-02")

	emailMetaData := dto.NewEmailMetaData(
		incomingInvoicePayload.Email,
		"",
		"",
		"html_templates/basic_template.html",
		incomingInvoicePayload.InvoiceID.String()+"-booking-details",
		eventData.EventName,
		strconv.Itoa(int(incomingInvoicePayload.SeatID)),
		incomingInvoicePayload.Email,
		formattedDate,
		incomingInvoicePayload.BookingID.String())

	//Check the status
	status := incomingInvoicePayload.PaymentStatus == datastruct.SUCCESS
	if !status {
		emailMetaData.EmailSubject = "Ticket App - Payment Failed"
		emailMetaData.BodyMessage = "Your payment failed, please make a new booking."
		err = emails.SendEmail(emailMetaData)
		if err != nil {
			response.ErrorResponse(w, http.StatusInternalServerError, "Fail to send email")
		}

		//Update seat status to booked if status success
		newSeatStatus := datastruct.Seat{
			Status: datastruct.OPEN,
		}

		// Retry attempt if update seat status is failed
		// Simple mechanism
		attemptCount := 3
		var updateSeatErr error

		for i := 0; i < attemptCount; i++ {
			_, updateSeatErr = m.seatService.UpdateSeat(incomingInvoicePayload.SeatID, newSeatStatus)
			if updateSeatErr == nil {
				break
			}
			fmt.Println("Update seat failed, retrying...")
			time.Sleep(time.Duration(math.Pow(2, float64(i))) * time.Second)
		}

		if updateSeatErr != nil {
			// Send email notification about the failure to update seat status
			emailMetaData.EmailSubject = "Ticket App - Seat Update Failed"
			emailMetaData.BodyMessage = "Failed to update seat status, please contact the administrator."
			err = emails.SendEmail(emailMetaData)
			if err != nil {
				fmt.Println("Fail to send email about seat update failure:", err)
			}
		}

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
	var updateSeatErr error

	for i := 0; i < attemptCount; i++ {
		_, updateSeatErr = m.seatService.UpdateSeat(incomingInvoicePayload.SeatID, newSeatStatus)
		if updateSeatErr == nil {
			break
		}
		fmt.Println("Update seat failed, retrying...")
		time.Sleep(time.Duration(math.Pow(2, float64(i))) * time.Second)
	}

	if updateSeatErr != nil {
		// Send email notification about the failure to update seat status
		emailMetaData.EmailSubject = "Ticket App - Seat Update Failed"
		emailMetaData.BodyMessage = "Failed to update seat status, please contact the administrator."
		err = emails.SendEmail(emailMetaData)
		if err != nil {
			fmt.Println("Fail to send email about seat update failure:", err)
		}
		return
	}

	//Send to user email?
	emailMetaData.EmailSubject = "Ticket App - Payment Success"
	emailMetaData.BodyMessage = "Payment is successful, here is the attached pdf for your booking"
	err = emails.SendEmail(emailMetaData)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, "Fail to send email")
	}

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
func (m *MicroserviceServer) WebhookCancelTicketHandler(w http.ResponseWriter, r *http.Request) {
	var incomingInvoicePayload dto.IncomingInvoicePayload
	err := json.NewDecoder(r.Body).Decode(&incomingInvoicePayload)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
		return
	}
	//Check the status
	status := incomingInvoicePayload.PaymentStatus == datastruct.SUCCESS
	fmt.Print(status)
	response.SuccessResponse(w, http.StatusOK, "Webhook is called", nil)
}
