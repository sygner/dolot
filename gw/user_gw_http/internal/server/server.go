package server

import (
	"dolott_user_gw_http/internal/client"
	"dolott_user_gw_http/internal/constants"
	"dolott_user_gw_http/internal/controllers"
	"dolott_user_gw_http/internal/middleware"
	"dolott_user_gw_http/internal/routes"
	"dolott_user_gw_http/internal/services"
	authentication_pb "dolott_user_gw_http/proto/api/authentication"
	game_pb "dolott_user_gw_http/proto/api/game"
	profile_pb "dolott_user_gw_http/proto/api/profile"
	ticket_pb "dolott_user_gw_http/proto/api/ticket"
	wallet_pb "dolott_user_gw_http/proto/api/wallet"
	"fmt"
	"log"
	"safir/libs/appconfigs"
	"safir/libs/appstates"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// RunServer sets up and starts the main server functionality.
func RunServer() {
	// Define variables for server configuration obtained from environment variables.
	var (
		listenAddress               = appconfigs.String("listen-address", "Server listen address")
		appDomain                   = appconfigs.String("app-domain", "App domain that use for api like a domain dolott.com enter valid ip if you don't have any domain or leave it empty")
		authenticationServerAddress = appconfigs.String("authentication-server-address", "Room server address")
		gameAddress                 = appconfigs.String("game-server-address", "Game server address")
		profileAddress              = appconfigs.String("profile-server-address", "Profile server address")
		walletAddress               = appconfigs.String("wallet-server-address", "Wallet server address")
		ticketAddress               = appconfigs.String("ticket-server-address", "Wallet server address")
		fileStoragePath             = appconfigs.String("file-storage-path", "File storage server address")
	)

	// Handle configuration errors and missing environment parameters.
	if err := appconfigs.Parse(); err != nil {
		appstates.PanicMissingEnvParams(err.Error())
	}
	if strings.Trim(*appDomain, "") == "" || len(*appDomain) <= 4 {
		appDomain = listenAddress
	}

	// Establish connections to external services via gRPC.
	var (
		authenticationServerConnection = client.GrpcClientServerConnection(*authenticationServerAddress)
		gameServerConnection           = client.GrpcClientServerConnection(*gameAddress)
		profileServerConnection        = client.GrpcClientServerConnection(*profileAddress)
		walletServiceConnection        = client.GrpcClientServerConnection(*walletAddress)
		ticketServiceConnection        = client.GrpcClientServerConnection(*ticketAddress)

		authenticationClient = authentication_pb.NewAuthentcationServiceClient(authenticationServerConnection)
		tokenClient          = authentication_pb.NewTokenServiceClient(authenticationServerConnection)
		userClient           = authentication_pb.NewUserServiceClient(authenticationServerConnection)

		winnerClient   = game_pb.NewWinnerServiceClient(gameServerConnection)
		gameClient     = game_pb.NewGameServiceClient(gameServerConnection)
		userGameClient = game_pb.NewUserServiceClient(gameServerConnection)

		profileClient = profile_pb.NewProfileServiceClient(profileServerConnection)

		walletClient = wallet_pb.NewWalletServiceClient(walletServiceConnection)

		ticketClient = ticket_pb.NewTicketServiceClient(ticketServiceConnection)
	)
	fmt.Println(*walletAddress)
	fmt.Println(walletClient)

	// Initialize different services used in the application.
	var (
		middleware middleware.MiddlewareService = middleware.NewMiddlewareService(tokenClient)

		authenticationService services.AuthenticationService = services.NewAuthenticationService(authenticationClient, userClient, tokenClient)
		userService           services.UserService           = services.NewUserService(userClient)

		winnerService   services.WinnerService   = services.NewWinnerService(winnerClient)
		gameService     services.GameService     = services.NewGameService(gameClient, ticketClient, walletClient, winnerClient, profileClient, *appDomain)
		userGameService services.UserGameService = services.NewUserGameService(userGameClient, ticketClient)
		walletService   services.WalletService   = services.NewWalletService(walletClient)
		ticketService   services.TicketService   = services.NewTicketService(ticketClient, walletClient)

		jobQueue services.JobQueue = services.NewJobQueue(gameService, walletService)

		profileService services.ProfileService = services.NewProfileService(profileClient, walletClient)

		authenticationController controllers.AuthenticationController = controllers.NewAuthenticationController(authenticationService, profileService, walletService)
		userController           controllers.UserController           = controllers.NewUserController(userService)
		profileController        controllers.ProfileController        = controllers.NewProfileController(profileService, walletService)

		winnerController   controllers.WinnerController = controllers.NewWinnerController(winnerService)
		gameController     controllers.GameController   = controllers.NewGameController(gameService, *fileStoragePath)
		userGameController controllers.UserGameHandler  = controllers.NewUserGameHandler(userGameService)

		walletController controllers.WalletController = controllers.NewWalletController(walletService)

		ticketController controllers.TicketController = controllers.NewTicketController(ticketService)

		adminController controllers.AdminController = controllers.NewAdminController()
	)
	// get the main account
	res, cerr := walletService.GetWalletsByUserId(0)
	if cerr != nil {
		log.Fatalf("failed to get main wallet: %v", cerr)
	}
	if len(res) == 0 {
		log.Fatalf("main wallet not found")
	}
	for _, wallet := range res {
		if wallet.CoinId == 1 {
			constants.MAIN_LUNC_WALLET_ID = wallet.Id
			constants.MAIN_LUNC_USER_WALLET_ID = wallet.UserId
			constants.MAIN_LUNC_WALLET_ADDRESS = wallet.Address
			break
		}
	}

	services.JOB_QUEUE = jobQueue
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
	routes.WalletGroup(v1, walletController, middleware)
	routes.TicketGroup(v1, ticketController, middleware)
	routes.AdminGroup(v1, adminController, middleware)

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
