package routes

import (
	"dolott_user_gw_http/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthenticationGroup(app fiber.Router, authenticationController controllers.AuthenticationController) {
	authenticationGroup := app.Group("/auth")
	authenticationGroup.Post("/signup", authenticationController.Signup)
	authenticationGroup.Post("/signin", authenticationController.Signin)
	authenticationGroup.Post("/verify", authenticationController.Verify)
	authenticationGroup.Post("/password/forgot", authenticationController.ForgotPassword)
	authenticationGroup.Post("/token/renew", authenticationController.RenewToken)
}
