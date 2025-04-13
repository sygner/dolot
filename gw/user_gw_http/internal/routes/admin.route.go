package routes

import (
	"dolott_user_gw_http/internal/controllers"
	"dolott_user_gw_http/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func AdminGroup(app fiber.Router, gameController controllers.AdminController, middleware middleware.MiddlewareService) {
	adminGroup := app.Group("/a1/admin")

	adminGroup.Use(middleware.VerificationMiddleware)
	adminGroup.Get("/rates", gameController.GetRates)
	adminGroup.Put("/rates", gameController.UpdateRates)
}
