package routes

import (
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/adapter/structs"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/app"
)

func PublicRoutes(server app.MicroserviceServer) structs.RoutePrefix {
	return structs.RoutePrefix{
		Prefix: "/v1/public",
		SubRoutes: []structs.Route{
			{
				"Send an email",
				"POST",
				"",
				server.SendEmailPublic,
				false,
			},
		},
	}
}
