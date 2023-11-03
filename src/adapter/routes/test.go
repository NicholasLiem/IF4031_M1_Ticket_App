package routes

import (
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/adapter/structs"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/app"
)

func TestRoutes(server app.MicroserviceServer) structs.RoutePrefix {
	return structs.RoutePrefix{
		Prefix: "/v1/test",
		SubRoutes: []structs.Route{
			{
				"Create a new user",
				"POST",
				"/",
				server.Test,
				false,
			},
		},
	}
}
