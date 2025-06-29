package server

import (
	"dolott_ticket/internal/database"
	"dolott_ticket/internal/handlers"
	"dolott_ticket/internal/repository"
	"dolott_ticket/internal/services"
	pb "dolott_ticket/proto/api"
	"log"
	"neo/libs/appconfigs"
	"neo/libs/appstates"
	"net"

	"google.golang.org/grpc"
)

func RunServer() {
	var (
		listenAddress = appconfigs.String("listen-address", "Server listen address")
		dbHost        = appconfigs.String("db-host", "Database host address") // Define a variable for the database host address.
		dbPort        = appconfigs.Int("db-port", "Database port")            // Define a variable for the database port.
		dbName        = appconfigs.String("db-name", "Database name")         // Define a variable for the database name.
		dbUsername    = appconfigs.String("db-username", "Database username") // Define a variable for the database username.
		dbPassword    = appconfigs.String("db-password", "Database password")
	)
	// Handle configuration errors.
	if err := appconfigs.Parse(); err != nil {
		appstates.PanicMissingEnvParams(err.Error()) // Log an error if there are missing environment parameters.
	}

	// Connect to the PostgreSQL database.
	db, err := database.ConnectToPostgres(*dbHost, *dbPort, *dbName, *dbUsername, *dbPassword)
	if err != nil {
		appstates.PanicDBConnectionFailed(err.Error()) // Log an error if the database connection fails.
	}

	var (
		repository repository.TicketRepository = repository.NewTicketRepository(db)

		ticketService services.TicketService = services.NewTicketService(repository)
		ticketHandler handlers.TicketHandler = *handlers.NewTicketHandler(ticketService)
	)

	listener, err := net.Listen("tcp", *listenAddress)
	if err != nil {
		log.Fatalf("error: %v", err)
		appstates.PanicServerSocketFailure(err.Error())
	}
	// Create a new gRPC server instance.
	grpcServer := grpc.NewServer()

	pb.RegisterTicketServiceServer(grpcServer, &ticketHandler)
	// Start serving the gRPC server.
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("error: %v", err)
	}

}
