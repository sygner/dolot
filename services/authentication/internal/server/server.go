package server

import (
	"dolott_authentication/internal/database"
	"dolott_authentication/internal/global"
	"dolott_authentication/internal/handlers"
	"dolott_authentication/internal/redis"
	"dolott_authentication/internal/repository"
	"dolott_authentication/internal/services"
	pb "dolott_authentication/proto/api"
	"log"
	"net"
	"safir/libs/appconfigs"
	"safir/libs/appstates"

	"google.golang.org/grpc"
)

func RunServer() {
	// Define variables for server configuration obtained from environment variables.
	var (
		listenAddress    = appconfigs.String("listen-address", "Server listen address")
		dbHost           = appconfigs.String("db-host", "Database host address") // Define a variable for the database host address.
		dbPort           = appconfigs.Int("db-port", "Database port")            // Define a variable for the database port.
		dbName           = appconfigs.String("db-name", "Database name")         // Define a variable for the database name.
		dbUsername       = appconfigs.String("db-username", "Database username") // Define a variable for the database username.
		dbPassword       = appconfigs.String("db-password", "Database password")
		redisHost        = appconfigs.String("rd-host", "Redis Host")
		templateFilePath = appconfigs.String("template-file-path", "Template File Path")
	)
	// Handle configuration errors.
	if err := appconfigs.Parse(); err != nil {
		appstates.PanicMissingEnvParams(err.Error()) // Log an error if there are missing environment parameters.
	}
	global.TEMPLATE_FILE_PATH = *templateFilePath

	// Connect to the PostgreSQL database.
	db, err := database.ConnectToPostgres(*dbHost, *dbPort, *dbName, *dbUsername, *dbPassword)
	if err != nil {
		appstates.PanicDBConnectionFailed(err.Error()) // Log an error if the database connection fails.
	}
	// Connect to the Redis.
	rd, err := redis.ConnectToRedis(*redisHost)
	if err != nil {
		appstates.PanicDBConnectionFailed(err.Error()) // Log an error if the database connection fails.
	}

	var (
		repository repository.AuthenticationRepository = repository.NewAuthenticationRepository(db)

		authenticationService services.AuthenticationService = services.NewAuthenticationService(repository, rd)
		userService           services.UserService           = services.NewUserService(repository, rd)
		tokenService          services.TokenService          = services.NewTokenService(repository)
		loginHistoryService   services.LoginHistoryService   = services.NewLoginHistoryService(repository)

		authenticationHandler handlers.AuthenticationHandler = *handlers.NewAuthenticationHandler(authenticationService)
		userHandler           handlers.UserHandler           = *handlers.NewUserHandler(userService, loginHistoryService)
		tokenHandler          handlers.TokenHandler          = *handlers.NewTokenHandler(tokenService)
	)

	listener, err := net.Listen("tcp", *listenAddress)
	if err != nil {
		log.Fatalf("error: %v", err)
		appstates.PanicServerSocketFailure(err.Error())
	}
	// Create a new gRPC server instance.
	grpcServer := grpc.NewServer()

	pb.RegisterAuthentcationServiceServer(grpcServer, &authenticationHandler)
	pb.RegisterUserServiceServer(grpcServer, &userHandler)
	pb.RegisterTokenServiceServer(grpcServer, &tokenHandler)
	// Start serving the gRPC server.
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("error: %v", err)
	}
}
