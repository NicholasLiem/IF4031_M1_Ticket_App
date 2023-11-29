package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/NicholasLiem/IF4031_M1_Ticket_App/adapter"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/adapter/clients"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/app"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/datastruct"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/repository"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/internal/service"
	"github.com/NicholasLiem/IF4031_M1_Ticket_App/seeder"
	"github.com/joho/godotenv"
)

func main() {

	/**
	Creating context
	*/
	ctx := context.Background()

	/**
	Loading .env file
	*/
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	/**
	Setting up http client
	*/
	clientAPIURL := os.Getenv("BASE_CLIENT_APP_URL")
	paymentAPIURL := os.Getenv("BASE_PAYMENT_APP_URL")
	apiIdentifierToken := os.Getenv("TICKET_API_KEY")
	headers := map[string]string{
		"Authorization": "Bearer " + apiIdentifierToken,
		"Content-Type":  "application/json",
	}
	restClientToClientApp := clients.NewRestClient(clientAPIURL, headers)
	restClientToPaymentApp := clients.NewRestClient(paymentAPIURL, headers)

	/**
	Setting up DB
	*/
	db := repository.SetupDB()
	redis := repository.SetupRedis(ctx)

	/**
	Registering DAO's and Services
	*/
	dao := repository.NewDAO(db, redis)

	eventService := service.NewEventService(dao)
	seatService := service.NewSeatService(dao)

	/**
	Registering Services to Server
	*/
	server := app.NewMicroservice(
		*restClientToClientApp,
		*restClientToPaymentApp,
		eventService,
		seatService,
	)

	/**
	Run DB Migration
	*/
	datastruct.Migrate(db, &datastruct.Event{}, &datastruct.Seat{})

	/**
	Seeder DB
	*/
	seedFlag := flag.Bool("seed", false, "Seed the database")
	flag.Parse()

	if *seedFlag {
		log.Println("Seeding the database...")
		seeder.Seed(db)
	}

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
