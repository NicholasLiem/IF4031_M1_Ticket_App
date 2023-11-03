package app

import "github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/service"

type MicroserviceServer struct {
	userService  service.UserService
	authService  service.AuthService
	eventService service.EventService
	seatService  service.SeatService
}

func NewMicroservice(
	userService service.UserService,
	authService service.AuthService,
	eventService service.EventService,
	seatService service.SeatService,
) *MicroserviceServer {
	return &MicroserviceServer{
		userService:  userService,
		authService:  authService,
		eventService: eventService,
		seatService:  seatService,
	}
}
