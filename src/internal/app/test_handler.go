package app

import (
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/dto"
	response "github.com/NicholasLiem/IF4031_M1_Ticket_App/utils/http"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/utils/messages"
	"net/http"
)

func (m *MicroserviceServer) Test(w http.ResponseWriter, r *http.Request) {
	var testDTO dto.TicketAppBookingResponseDTO

	testDTO = dto.TicketAppBookingResponseDTO{
		BookingID: 1,
		Status:    dto.BookingOnProcess,
		Message:   "Berhasil",
	}

	response.SuccessResponse(w, http.StatusOK, messages.SuccessfulDataObtain, testDTO)
	return
}
