package routes

import (
	"dolott_user_gw_http/internal/controllers"
	"dolott_user_gw_http/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserGroup(app fiber.Router, userController controllers.UserController, middleware middleware.MiddlewareService) {
	userGroup := app.Group("/usr")
	userGroup.Use(middleware.VerificationMiddleware)
	userGroup.Get("/u/:user_id", userController.GetUserByUserId)
	userGroup.Get("/s", userController.GetSelfData)
	userGroup.Post("/e", userController.GetUserByEmail)
	userGroup.Post("/u", userController.GetUserAccountUsername)
	userGroup.Post("/login/history", userController.GetLoginHistoryByUserId)
	userGroup.Post("/password/reset", userController.ResetPassword)
}
