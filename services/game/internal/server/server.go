package server

import (
	"dolott_game/internal/database"
	"dolott_game/internal/handlers"
	"dolott_game/internal/repository"
	"dolott_game/internal/services"
	pb "dolott_game/proto/api"
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
		repository repository.GameRepository = repository.NewGameRepository(db)

		gameService   services.GameServices   = services.NewGameServices(repository)
		userService   services.UserServices   = services.NewUserServices(repository)
		winnerService services.WinnerServices = services.NewWinnerServices(repository)

		gameHandler   handlers.GameHandler   = *handlers.NewGameHandler(gameService)
		userHandler   handlers.UserHandler   = *handlers.NewUserHandler(userService)
		winnerHandler handlers.WinnerHandler = *handlers.NewWinenrHandler(winnerService)
	)

	listener, err := net.Listen("tcp", *listenAddress)
	if err != nil {
		log.Fatalf("error: %v", err)
		appstates.PanicServerSocketFailure(err.Error())
	}
	// Create a new gRPC server instance.
	grpcServer := grpc.NewServer()

	pb.RegisterGameServiceServer(grpcServer, &gameHandler)
	pb.RegisterUserServiceServer(grpcServer, &userHandler)
	pb.RegisterWinnerServiceServer(grpcServer, &winnerHandler)
	// Start serving the gRPC server.
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("error: %v", err)
	}

}
