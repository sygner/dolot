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

}
