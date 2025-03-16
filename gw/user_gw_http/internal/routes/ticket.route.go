package routes

import (
	"dolott_user_gw_http/internal/controllers"
	"dolott_user_gw_http/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func TicketGroup(app fiber.Router, ticketController controllers.TicketController, middleware middleware.MiddlewareService) {
	ticketGroup := app.Group("/ticket")
	ticketGroup.Use(middleware.VerificationMiddleware)
	ticketGroup.Post("/single/add", ticketController.AddTicket)
	ticketGroup.Get("/signature/:signature", ticketController.GetTicketBySignatureAndUserId)
	ticketGroup.Get("/id/:ticket_id", ticketController.GetTicketByUserIdAndTicketId)
	ticketGroup.Post("/all", ticketController.GetAllUserTickets)
	ticketGroup.Post("/open/all", ticketController.GetUserOpenTickets)
	ticketGroup.Post("/used/all", ticketController.GetAllUsedTickets)
	ticketGroup.Post("/use", ticketController.UseTickets)
	ticketGroup.Post("/multiple/all", ticketController.AddTickets)
	ticketGroup.Get("/game/:game_id", ticketController.GetAllUserTicketsByGameId)
	ticketGroup.Get("/u/game/:game_id", ticketController.GetAllTicketsByGameId)
	ticketGroup.Post("/buy", ticketController.BuyTickets)

}
