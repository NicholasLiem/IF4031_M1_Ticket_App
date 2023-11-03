package routes

import (
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/adapter/structs"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/app"
)

func SeatRoutes(server app.MicroserviceServer) structs.RoutePrefix {
	return structs.RoutePrefix{
		Prefix: "/v1/seat",
		SubRoutes: []structs.Route{
			{
				"Create a new seat",
				"POST",
				"/{event_id}",
				server.CreateSeat,
				true,
			},
			{
				"Get a seat by id",
				"GET",
				"/{seat_id}",
				server.GetSeat,
				true,
			},
			{
				"Update a seat by id",
				"PUT",
				"/{seat_id}",
				server.UpdateSeat,
				true,
			},
			{
				"Delete a seat",
				"DELETE",
				"/{seat_id}",
				server.DeleteSeat,
				true,
			},
		},
	}
}
