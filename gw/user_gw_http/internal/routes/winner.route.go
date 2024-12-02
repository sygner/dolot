package routes

import (
	"dolott_user_gw_http/internal/controllers"
	"dolott_user_gw_http/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func WinnerGroup(app fiber.Router, winnerController controllers.WinnerHandler, middleware middleware.MiddlewareService) {
	winnerGroup := app.Group("/game/winner")
	winnerGroup.Use(middleware.VerificationMiddleware)
	winnerGroup.Get("/:game_id", winnerController.GetWinnersByGameId)
}
