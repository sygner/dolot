package controllers

import (
	"dolott_user_gw_http/internal/admin"
	"dolott_user_gw_http/internal/models"

	"github.com/gofiber/fiber/v2"
)

type (
	AdminController interface {
		GetRates(*fiber.Ctx) error
		UpdateRates(*fiber.Ctx) error
	}
	adminController struct {
	}
)

func NewAdminController() AdminController {
	return &adminController{}
}

func (c *adminController) GetRates(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]interface{}{
		"ticket_buy_rate":            admin.TICKET_BUY_RATE,
		"transaction_tax_percentage": admin.TRANSACTION_TAX_PERCENTAGE,
		"impression_exchange_rate":   admin.IMPRESSION_EXCHANGE_RATE,
		"success":                    true,
	})
}

func (c *adminController) UpdateRates(ctx *fiber.Ctx) error {
	updateRatesDTO := new(models.UpdateRatesDTO)
	if err := ctx.BodyParser(updateRatesDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #189",
			"success": false,
		})
	}

	if updateRatesDTO.TicketBuyRate != nil {
		admin.TICKET_BUY_RATE = *updateRatesDTO.TicketBuyRate
	}
	if updateRatesDTO.TransactionTaxPercentage != nil {
		admin.TRANSACTION_TAX_PERCENTAGE = *updateRatesDTO.TransactionTaxPercentage
	}
	if updateRatesDTO.ImpressionExchangeRate != nil {
		admin.IMPRESSION_EXCHANGE_RATE = *updateRatesDTO.ImpressionExchangeRate
	}

	return ctx.Status(200).JSON(map[string]interface{}{
		"message": "done",
		"success": false,
	})
}
