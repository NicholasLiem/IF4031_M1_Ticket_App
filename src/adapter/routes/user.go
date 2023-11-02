package routes

import (
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/adapter/structs"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/app"
)

func UserRoutes(server app.MicroserviceServer) structs.RoutePrefix {
	return structs.RoutePrefix{
		Prefix: "/v1/user",
		SubRoutes: []structs.Route{
			{
				"Create a new user",
				"POST",
				"/{user_id}",
				server.CreateUser,
				true,
			},
			{
				"Update a user",
				"PUT",
				"/{user_id}",
				server.UpdateUser,
				true,
			},
			{
				"Delete a user",
				"DELETE",
				"/{user_id}",
				server.DeleteUser,
				true,
			},
			{
				"Get user data",
				"GET",
				"/{user_id}",
				server.GetUserData,
				true,
			},
		},
	}
}
