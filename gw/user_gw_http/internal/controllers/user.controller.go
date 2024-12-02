package controllers

import (
	"dolott_user_gw_http/internal/models"
	"dolott_user_gw_http/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type (
	UserController interface {
		GetUserByUserId(*fiber.Ctx) error
		GetUserByEmail(*fiber.Ctx) error
		GetUserAccountUsername(*fiber.Ctx) error
		GetSelfData(*fiber.Ctx) error
		GetLoginHistoryByUserId(*fiber.Ctx) error
		ResetPassword(*fiber.Ctx) error
	}
	userController struct {
		userService services.UserService
	}
)

func NewUserController(userService services.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (c *userController) GetUserByUserId(ctx *fiber.Ctx) error {
	userIdS := ctx.Params("user_id")
	if userIdS == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "user id cannot be empty, error code #103",
			"success": false,
		})
	}
	userId, err := strconv.ParseInt(userIdS, 10, 64)
	if err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "bad request, error code #104",
			"success": false,
		})
	}
	res, rerr := c.userService.GetUserByUserId(int32(userId))
	if rerr != nil {
		return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *userController) GetUserByEmail(ctx *fiber.Ctx) error {
	emailDTO := new(models.EmailDTO)
	if err := ctx.BodyParser(emailDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #105",
			"success": false,
		})
	}
	res, rerr := c.userService.GetUserByEmail(emailDTO.Email)
	if rerr != nil {
		return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *userController) GetUserAccountUsername(ctx *fiber.Ctx) error {
	accountUsernameDTO := new(models.AccountUsernameDTO)
	if err := ctx.BodyParser(accountUsernameDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #105",
			"success": false,
		})
	}
	res, rerr := c.userService.GetUserByAccountUsername(accountUsernameDTO.AccountUsername)
	if rerr != nil {
		return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *userController) GetSelfData(ctx *fiber.Ctx) error {
	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #106",
			"success": false,
		})
	}
	userData := localData.(models.Token)
	res, rerr := c.userService.GetUserByUserId(userData.UserId)
	if rerr != nil {
		return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *userController) GetLoginHistoryByUserId(ctx *fiber.Ctx) error {
	paginationDTO := new(models.Pagination)
	if err := ctx.BodyParser(paginationDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #107",
			"success": false,
		})
	}

	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #108",
			"success": false,
		})
	}
	userData := localData.(models.Token)
	res, err := c.userService.GetLoginHistoryByUserId(paginationDTO, userData.UserId)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *userController) ResetPassword(ctx *fiber.Ctx) error {
	resetPasswordDTO := new(models.ResetPasswordDTO)
	if err := ctx.BodyParser(resetPasswordDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #109",
			"success": false,
		})
	}

	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #110",
			"success": false,
		})
	}
	userData := localData.(models.Token)
	resetPasswordDTO.UserId = userData.UserId
	err := c.userService.ResetPassword(resetPasswordDTO)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"message": "password changed",
		"success": true,
	})
}
