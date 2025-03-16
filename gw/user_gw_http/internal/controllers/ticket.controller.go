package controllers

import (
	"dolott_user_gw_http/internal/models"
	"dolott_user_gw_http/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type (
	TicketController interface {
		AddTicket(*fiber.Ctx) error
		GetTicketBySignatureAndUserId(*fiber.Ctx) error
		GetTicketByUserIdAndTicketId(*fiber.Ctx) error
		GetAllUserTickets(*fiber.Ctx) error
		GetUserOpenTickets(*fiber.Ctx) error
		GetAllUsedTickets(*fiber.Ctx) error
		UseTickets(*fiber.Ctx) error
		AddTickets(*fiber.Ctx) error
		GetAllUserTicketsByGameId(*fiber.Ctx) error
		GetAllTicketsByGameId(*fiber.Ctx) error
		BuyTickets(*fiber.Ctx) error
	}
	ticketController struct {
		ticketService services.TicketService
	}
)

func NewTicketController(ticketService services.TicketService) TicketController {
	return &ticketController{
		ticketService: ticketService,
	}
}

func (c *ticketController) AddTicket(ctx *fiber.Ctx) error {
	addTicketDTO := new(models.AddTicketDTO)
	if err := ctx.BodyParser(addTicketDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #154",
			"success": false,
		})
	}

	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #155",
			"success": false,
		})
	}
	userData := localData.(models.Token)
	addTicketDTO.UserId = userData.UserId
	// addTicketDTO.TicketType = "testing"

	res, err := c.ticketService.AddTicket(addTicketDTO)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *ticketController) GetTicketBySignatureAndUserId(ctx *fiber.Ctx) error {
	signature := ctx.Params("signature")
	if signature == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "signature cannot be empty, error code #156",
			"success": false,
		})
	}

	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #157",
			"success": false,
		})
	}
	userData := localData.(models.Token)
	res, err := c.ticketService.GetTicketBySignatureAndUserId(userData.UserId, signature)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *ticketController) GetTicketByUserIdAndTicketId(ctx *fiber.Ctx) error {
	ticketIdString := ctx.Params("ticket_id")
	if ticketIdString == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "ticket id cannot be empty, error code #158",
			"success": false,
		})
	}

	ticketIdInt, rerr := strconv.Atoi(ticketIdString)
	if rerr != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "invalid ticket id, must be a number, error code #159",
			"success": false,
		})
	}

	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #160",
			"success": false,
		})
	}
	userData := localData.(models.Token)

	res, err := c.ticketService.GetTicketByUserIdAndTicketId(userData.UserId, int32(ticketIdInt))

	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *ticketController) GetAllUserTickets(ctx *fiber.Ctx) error {
	paginationDTO := new(models.Pagination)
	if err := ctx.BodyParser(paginationDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #161",
			"success": false,
		})
	}
	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #162",
			"success": false,
		})
	}
	userData := localData.(models.Token)

	res, err := c.ticketService.GetAllUserTickets(userData.UserId, paginationDTO)

	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *ticketController) GetUserOpenTickets(ctx *fiber.Ctx) error {
	paginationDTO := new(models.Pagination)
	if err := ctx.BodyParser(paginationDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #163",
			"success": false,
		})
	}
	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #164",
			"success": false,
		})
	}
	userData := localData.(models.Token)

	res, err := c.ticketService.GetUserOpenTickets(userData.UserId, paginationDTO)

	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *ticketController) GetAllUsedTickets(ctx *fiber.Ctx) error {
	paginationDTO := new(models.Pagination)
	if err := ctx.BodyParser(paginationDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #165",
			"success": false,
		})
	}
	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #166",
			"success": false,
		})
	}
	userData := localData.(models.Token)

	res, err := c.ticketService.GetAllUsedTickets(userData.UserId, paginationDTO)

	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *ticketController) UseTickets(ctx *fiber.Ctx) error {
	useTicketDTO := new(models.UseTicketDTO)
	if err := ctx.BodyParser(useTicketDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #167",
			"success": false,
		})
	}

	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #168",
			"success": false,
		})
	}
	userData := localData.(models.Token)

	res, err := c.ticketService.UseTickets(userData.UserId, useTicketDTO.Amount, useTicketDTO.GameId)

	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *ticketController) AddTickets(ctx *fiber.Ctx) error {
	addTicketsDTO := new(models.AddTicketsDTO)
	if err := ctx.BodyParser(addTicketsDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #169",
			"success": false,
		})
	}

	shouldReturn := ctx.QueryBool("return", false)
	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #170",
			"success": false,
		})
	}
	userData := localData.(models.Token)

	for i := range addTicketsDTO.AddTickets {
		addTicketsDTO.AddTickets[i].UserId = userData.UserId
		// addTicketsDTO.AddTickets[i].TicketType = "testing"
	}

	res, err := c.ticketService.AddTickets(addTicketsDTO.AddTickets, shouldReturn)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *ticketController) GetAllUserTicketsByGameId(ctx *fiber.Ctx) error {
	gameId := ctx.Params("game_id")
	if gameId == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "game id cannot be empty, error code #171",
			"success": false,
		})
	}
	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #172",
			"success": false,
		})
	}
	userData := localData.(models.Token)

	res, err := c.ticketService.GetAllUserTicketsByGameId(userData.UserId, gameId)

	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *ticketController) GetAllTicketsByGameId(ctx *fiber.Ctx) error {
	gameId := ctx.Params("game_id")
	if gameId == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "game id cannot be empty, error code #173",
			"success": false,
		})
	}

	res, err := c.ticketService.GetAllTicketsByGameId(gameId)

	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *ticketController) BuyTickets(ctx *fiber.Ctx) error {
	buyTicketsDTO := new(models.BuyTicketDTO)
	if err := ctx.BodyParser(buyTicketsDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #174",
			"success": false,
		})
	}
	shouldReturn := ctx.QueryBool("return", false)

	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #175",
			"success": false,
		})
	}
	userData := localData.(models.Token)
	res, err := c.ticketService.BuyTickets(userData.UserId, buyTicketsDTO.TotalTickets, buyTicketsDTO.WalletId, shouldReturn)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	if res != nil && len(res.Tickets) > 0 {
		return ctx.JSON(map[string]interface{}{
			"data":    res,
			"success": true,
		})
	} else {
		return ctx.JSON(map[string]interface{}{
			"message": "tickets bought",
			"success": true,
		})
	}
}
