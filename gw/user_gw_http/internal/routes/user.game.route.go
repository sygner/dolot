package routes

import (
	"dolott_user_gw_http/internal/controllers"
	"dolott_user_gw_http/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserGameGroup(app fiber.Router, userGameHandlerController controllers.UserGameHandler, middleware middleware.MiddlewareService) {
	userGameGroup := app.Group("/game/usr")
	userGameGroup.Use(middleware.VerificationMiddleware)
	userGameGroup.Post("/add", userGameHandlerController.AddUserChoice)
	userGameGroup.Post("/multi/add", userGameHandlerController.AddUserChoices)
	userGameGroup.Post("/u", userGameHandlerController.GetUserChoicesByUserId)
	userGameGroup.Post("/t", userGameHandlerController.GetUserChoicesByUserIdAndTimeRange)
	userGameGroup.Post("/g/:game_id", userGameHandlerController.GetUserChoicesByGameIdAndPagination)
	userGameGroup.Get("/interaction/ids", userGameHandlerController.GetAllUserGames)
}
