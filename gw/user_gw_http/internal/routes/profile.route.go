package routes

import (
	"dolott_user_gw_http/internal/controllers"
	"dolott_user_gw_http/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProfileGroup(app fiber.Router, profileController controllers.ProfileController, middleware middleware.MiddlewareService) {
	profileGroup := app.Group("/profile")
	profileGroup.Use(middleware.VerificationMiddleware)
	profileGroup.Get("/usr/:username", profileController.GetProfileByUsername)
	profileGroup.Get("/self", profileController.GetSelfProfile)
	profileGroup.Put("/update", profileController.UpdateProfile)
	profileGroup.Get("/sid/:sid", profileController.GetProfileBySid)
	profileGroup.Put("/update/impression", profileController.ChangeUserImpression)
	profileGroup.Post("/impression/exchange", profileController.ImpressionExchange)
	profileGroup.Get("/search/username", profileController.SearchUsername)
	profileGroup.Get("/rank/usr", profileController.GetAllUserRanking)
	profileGroup.Get("/rank/leaderboard", profileController.GetUserLeaderBoard)
	profileGroup.Put("/rank/update", profileController.UpdateUserRank)
	profileGroup.Put("/update/impression/dcredit", profileController.ChangeImpressionAndDCoin)
}
