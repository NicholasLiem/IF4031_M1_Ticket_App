package routes

import (
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/adapter/structs"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/app"
)

func WebhookRoutes(server app.MicroserviceServer) structs.RoutePrefix {
	return structs.RoutePrefix{
		Prefix: "/v1/webhook",
		SubRoutes: []structs.Route{
			{
				"Incoming webhook from payment app",
				"POST",
				"",
				server.WebhookPaymentHandler,
				true,
			},
			{
				"Incoming webhook from payment app",
				"POST",
				"cancel",
				server.WebhookCancelTicketHandler,
				true,
			},
		},
	}
}
