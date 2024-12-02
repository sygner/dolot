package server

import (
	"dolott_profile/internal/database"
	"dolott_profile/internal/handlers"
	"dolott_profile/internal/repository"
	"dolott_profile/internal/services"
	pb "dolott_profile/proto/api"
	"log"
	"net"
	"safir/libs/appconfigs"
	"safir/libs/appstates"

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
		repository repository.ProfileRepository = repository.NewProfileRepository(db)

		profileService services.ProfileService = services.NewProfileRepository(repository)

		profileHandler handlers.ProfileHandler = *handlers.NewProfileHandler(profileService)
	)

	listener, err := net.Listen("tcp", *listenAddress)
	if err != nil {
		log.Fatalf("error: %v", err)
		appstates.PanicServerSocketFailure(err.Error())
	}
	// Create a new gRPC server instance.
	grpcServer := grpc.NewServer()

	pb.RegisterProfileServiceServer(grpcServer, &profileHandler)
	// Start serving the gRPC server.
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("error: %v", err)
	}

}
