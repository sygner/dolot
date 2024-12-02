package controllers

import (
	"dolott_user_gw_http/internal/models"
	"dolott_user_gw_http/internal/services"

	"github.com/gofiber/fiber/v2"
)

type (
	ProfileController interface {
		GetProfileByUsername(*fiber.Ctx) error
		GetSelfProfile(*fiber.Ctx) error
		UpdateProfile(*fiber.Ctx) error
		GetProfileBySid(*fiber.Ctx) error
	}
	profileController struct {
		profileService services.ProfileService
	}
)

func NewProfileController(profileService services.ProfileService) ProfileController {
	return &profileController{
		profileService: profileService,
	}
}

func (c *profileController) GetProfileByUsername(ctx *fiber.Ctx) error {
	username := ctx.Params("username")
	if username == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "username cannot be empty, error code #128",
			"success": false,
		})
	}
	res, err := c.profileService.GetProfileUsername(username)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}
func (c *profileController) GetProfileBySid(ctx *fiber.Ctx) error {
	sid := ctx.Params("sid")
	if sid == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "sid cannot be empty, error code #132",
			"success": false,
		})
	}
	res, err := c.profileService.GetProfileSid(sid)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *profileController) GetSelfProfile(ctx *fiber.Ctx) error {
	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #129",
			"success": false,
		})
	}
	userData := localData.(models.Token)
	res, err := c.profileService.GetProfileByUserId(userData.UserId)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *profileController) UpdateProfile(ctx *fiber.Ctx) error {
	updateProfileDTO := new(models.UpdateProfileDTO)
	if err := ctx.BodyParser(updateProfileDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #130",
			"success": false,
		})
	}
	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "internal issue, error code #131",
			"success": false,
		})
	}
	userData := localData.(models.Token)
	err := c.profileService.UpdateProfile(userData.UserId, updateProfileDTO.Username)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"message": "profile updated",
		"success": true,
	})
}
