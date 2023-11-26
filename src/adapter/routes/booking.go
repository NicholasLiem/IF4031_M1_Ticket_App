package routes

import (
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/adapter/structs"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/app"
)

func BookingRoutes(server app.MicroserviceServer) structs.RoutePrefix {
	return structs.RoutePrefix{
		Prefix: "/v1/book",
		SubRoutes: []structs.Route{
			{
				"Book seat",
				"POST",
				"",
				server.BookSeat,
				true,
			},
			{
				"Cancel Seat",
				"POST",
				"/cancel",
				server.CancelSeat,
				true,
			},
		},
	}
}
