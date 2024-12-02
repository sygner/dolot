package server

import (
	"dolott_user_gw_http/internal/client"
	"dolott_user_gw_http/internal/controllers"
	"dolott_user_gw_http/internal/middleware"
	"dolott_user_gw_http/internal/routes"
	"dolott_user_gw_http/internal/services"
	authentication_pb "dolott_user_gw_http/proto/api/authentication"
	game_pb "dolott_user_gw_http/proto/api/game"
	profile_pb "dolott_user_gw_http/proto/api/profile"
	"fmt"
	"safir/libs/appconfigs"
	"safir/libs/appstates"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// RunServer sets up and starts the main server functionality.
func RunServer() {
	// Define variables for server configuration obtained from environment variables.
	var (
		listenAddress               = appconfigs.String("listen-address", "Server listen address")
		authenticationServerAddress = appconfigs.String("authentication-server-address", "Room server address")
		gameAddress                 = appconfigs.String("game-server-address", "Game server address")
		profileAddress              = appconfigs.String("profile-server-address", "Profile server address")
	)

	// Handle configuration errors and missing environment parameters.
	if err := appconfigs.Parse(); err != nil {
		appstates.PanicMissingEnvParams(err.Error())
	}

	// Establish connections to external services via gRPC.
	var (
		authenticationServerConnection = client.GrpcClientServerConnection(*authenticationServerAddress)
		gameServerConnection           = client.GrpcClientServerConnection(*gameAddress)
		profileServerConnection        = client.GrpcClientServerConnection(*profileAddress)

		authenticationClient = authentication_pb.NewAuthentcationServiceClient(authenticationServerConnection)
		tokenClient          = authentication_pb.NewTokenServiceClient(authenticationServerConnection)
		userClient           = authentication_pb.NewUserServiceClient(authenticationServerConnection)

		winnerClient   = game_pb.NewWinnerServiceClient(gameServerConnection)
		gameClient     = game_pb.NewGameServiceClient(gameServerConnection)
		userGameClient = game_pb.NewUserServiceClient(gameServerConnection)

		profileClient = profile_pb.NewProfileServiceClient(profileServerConnection)
	)

	// Initialize different services used in the application.
	var (
		middleware middleware.MiddlewareService = middleware.NewMiddlewareService(tokenClient)

		authenticationService services.AuthenticationService = services.NewAuthenticationService(authenticationClient, userClient, tokenClient)
		userService           services.UserService           = services.NewUserService(userClient)

		winnerService   services.WinnerService   = services.NewWinnerService(winnerClient)
		gameService     services.GameService     = services.NewGameService(gameClient)
		userGameService services.UserGameService = services.NewUserGameService(userGameClient)

		profileService services.ProfileService = services.NewProfileService(profileClient)

		authenticationController controllers.AuthenticationController = controllers.NewAuthenticationController(authenticationService, profileService)
		userController           controllers.UserController           = controllers.NewUserController(userService)
		profileController        controllers.ProfileController        = controllers.NewProfileController(profileService)

		winnerController   controllers.WinnerHandler   = controllers.NewWinnerHandler(winnerService)
		gameController     controllers.GameHandler     = controllers.NewGameHandler(gameService)
		userGameController controllers.UserGameHandler = controllers.NewUserGameHandler(userGameService)
	)

	// Create a new Fiber instance.
	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// Use CORS middleware for handling cross-origin requests.
	app.Use(cors.New())

	v1 := app.Group("/api")
	// this route doesn't need any middleware to check jwt
	routes.AuthenticationGroup(v1, authenticationController)

	// routes they need middleware
	routes.WinnerGroup(v1, winnerController, middleware)
	routes.GameGroup(v1, gameController, middleware)
	routes.UserGroup(v1, userController, middleware)
	routes.UserGameGroup(v1, userGameController, middleware)

	routes.ProfileGroup(v1, profileController, middleware)

	// Default route for a simple hello world response.
	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println(c.Context())
		return c.SendString("Pong!")
	})

	// Start the Fiber server and log any errors encountered during startup.
	err := app.Listen(*listenAddress)
	if err != nil {
		fmt.Println(err)
	}
}
