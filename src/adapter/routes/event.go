package routes

import (
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/adapter/structs"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/app"
)

func EventRoutes(server app.MicroserviceServer) structs.RoutePrefix {
	return structs.RoutePrefix{
		Prefix: "/v1/event",
		SubRoutes: []structs.Route{
			{
				"Create a new event",
				"POST",
				"",
				server.CreateEvent,
				true,
			},
			{
				"Get an event by id",
				"GET",
				"/{event_id}",
				server.GetEvent,
				true,
			},
			{
				"Update an event by id",
				"PUT",
				"/{event_id}",
				server.UpdateEvent,
				true,
			},
			{
				"Delete an event",
				"DELETE",
				"/{event_id}",
				server.DeleteEvent,
				true,
			},
		},
	}
}
