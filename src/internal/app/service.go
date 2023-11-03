package app

import (
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/adapter/clients"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/service"
)

type MicroserviceServer struct {
	eventService           service.EventService
	seatService            service.SeatService
	restClientToClientApp  clients.RestClient
	restClientToPaymentApp clients.RestClient
}

func NewMicroservice(
	restClientToClientApp clients.RestClient,
	restClientToPaymentApp clients.RestClient,
	eventService service.EventService,
	seatService service.SeatService,
) *MicroserviceServer {
	return &MicroserviceServer{
		restClientToClientApp:  restClientToClientApp,
		restClientToPaymentApp: restClientToPaymentApp,
		eventService:           eventService,
		seatService:            seatService,
	}
}
