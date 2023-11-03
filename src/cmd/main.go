package main

import (
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/adapter"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/app"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/datastruct"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/repository"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/service"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {

	/**
	Creating context
	*/
	//ctx := context.Background()

	/**
	Loading .env file
	*/
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	/**
	Setting up DB
	*/
	db := repository.SetupDB()

	/**
	Registering DAO's and Services
	*/
	dao := repository.NewDAO(db)
	userService := service.NewUserService(dao)
	authService := service.NewAuthService(dao)
	eventService := service.NewEventService(dao)

	/**
	Registering Services to Server
	*/
	server := app.NewMicroservice(
		userService,
		authService,
		eventService,
	)

	/**
	Run DB Migration
	*/
	datastruct.Migrate(db, &datastruct.Event{}, &datastruct.UserModel{}, &datastruct.Payment{}, &datastruct.Invoice{}, &datastruct.Booking{})

	/**
	Setting up the router
	*/
	serverRouter := adapter.NewRouter(*server)

	/**
	Running the server
	*/
	port := os.Getenv("PORT")
	log.Println("Running the server on port " + port)

	if os.Getenv("ENVIRONMENT") == "DEV" {
		log.Fatal(http.ListenAndServe("127.0.0.1:"+port, serverRouter))
	}
	log.Fatal(http.ListenAndServe(":"+port, serverRouter))
}
