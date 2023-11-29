package app

import (
	"fmt"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/dto"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/utils/emails"
	response "github.com/NicholasLiem/IF4031_M1_Ticket_App/utils/http"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/utils/messages"
	"net/http"
)

func (m *MicroserviceServer) SendEmailPublic(w http.ResponseWriter, r *http.Request) {

	emailMetaData := dto.NewEmailMetaData(
		"nicholasliem01@gmail.com",
		"Email Subject",
		"This is the body message",
		"html_templates/basic_template.html",
		"booking-details",
		"Event Name",
		"Seat Id",
		"User Email",
		"User",
		"booking-qr",
	)

	err := emails.SendEmail(emailMetaData)
	if err != nil {
		fmt.Println(err)
		response.ErrorResponse(w, http.StatusInternalServerError, "Fail to send email")
		return
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulDataCreation, nil)
	return
}
