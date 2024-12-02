package routes

import (
	"dolott_user_gw_http/internal/controllers"
	"dolott_user_gw_http/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func GameGroup(app fiber.Router, gameController controllers.GameHandler, middleware middleware.MiddlewareService) {
	gameGroup := app.Group("/game")
	gameGroup.Use(middleware.VerificationMiddleware)
	gameGroup.Get("/get/:game_id", gameController.GetGameByGameId)
	gameGroup.Post("/add", gameController.AddGame)
	gameGroup.Get("/next/", gameController.GetNextGamesByGameType)
	gameGroup.Get("/next/games", gameController.GetAllNextGames)
	gameGroup.Get("/previous/games", gameController.GetAllPreviousGames)
	gameGroup.Get("/all", gameController.GetAllGames)
	gameGroup.Delete("/del/:game_id", gameController.DeleteGameByGameId)
	gameGroup.Get("/check/:game_id", gameController.CheckGameExistsByGameId)
	gameGroup.Post("/creator", gameController.GetGamesByCreatorId)
	gameGroup.Post("/result", gameController.AddResultByGameId)
}
