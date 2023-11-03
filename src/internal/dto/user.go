package dto

import "github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/datastruct"

type CreateUserDTO struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

type UpdateUserDTO struct {
	UserID    uint            `json:"user_id,omitempty"`
	Email     string          `json:"email,omitempty"`
	FirstName string          `json:"first_name,omitempty"`
	LastName  string          `json:"last_name,omitempty"`
	Password  string          `json:"password,omitempty"`
	Role      datastruct.Role `json:"role,omitempty"`
}
