package controllers

import (
	"dolott_user_gw_http/internal/models"
	"dolott_user_gw_http/internal/services"
	"dolott_user_gw_http/internal/types"
	"fmt"
	"neo/libs/idgen"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type (
	AuthenticationController interface {
		Signup(*fiber.Ctx) error
		Signin(*fiber.Ctx) error
		Verify(*fiber.Ctx) error
		ForgotPassword(*fiber.Ctx) error
		RenewToken(*fiber.Ctx) error
	}
	authenticationController struct {
		authenticationService services.AuthenticationService
		profileService        services.ProfileService
		walletService         services.WalletService
	}
)

func NewAuthenticationController(authenticationService services.AuthenticationService, profileService services.ProfileService, walletService services.WalletService) AuthenticationController {
	return &authenticationController{
		authenticationService: authenticationService,
		profileService:        profileService,
		walletService:         walletService,
	}
}

func (c *authenticationController) GenerateUsername(signupDTO *models.SignupDTO) (string, *types.Error) {
	var username string

	if signupDTO.Username != nil && *signupDTO.Username != "" {
		return *signupDTO.Username, nil // Return existing username if provided
	}

	// Generate a username based on the email
	emailParts := strings.Split(signupDTO.Email, "@")
	if len(emailParts) > 1 {
		username = emailParts[0]
		if len(username) < 6 {
			randomUsername, err := c.CreateRandomUsername(6)
			if err != nil {
				return "", err
			}
			username += "_" + randomUsername
		}
	} else {
		// Fully random username generation
		randomUsername, err := c.CreateRandomUsername(12)
		if err != nil {
			return "", err
		}
		username = signupDTO.Email + "_" + randomUsername
	}

	// Ensure the username is unique
	for i := 0; i < 20; i++ {
		exists, err := c.profileService.CheckUsernameExists(username)
		if err != nil {
			return "", err
		}
		if !exists {
			return username, nil
		}

		// Generate a new random username if the current one exists
		randomUsername, err := c.CreateRandomUsername(6)
		if err != nil {
			return "", err
		}
		username += "_" + randomUsername
	}

	return "", types.NewAlreadyExistsError("failed to generate a unique username after multiple attempts")
}

func (c *authenticationController) Signup(ctx *fiber.Ctx) error {
	// Parse request body into SignupDTO
	signupDTO := new(models.SignupDTO)
	if err := ctx.BodyParser(signupDTO); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Invalid body, error code #100",
			"success": false,
		})
	}

	signupDTO.Ip = ctx.IP()
	signupDTO.Agent = string(ctx.Context().UserAgent())

	res, err := c.authenticationService.Signup(signupDTO)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	if res != nil {
		profile, rerr := c.profileService.GetProfileByUserId(res.UserId)
		wallets := make([]models.Wallet, 0)
		if rerr != nil {
			if rerr.Code != 5 {
				return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
			}

			username, err := c.GenerateUsername(signupDTO)
			if err != nil {
				return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
			}

			profile, rerr = c.profileService.AddProfile(res.UserId, username)
			if rerr != nil {
				return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
			}
			wallet, rerr := c.walletService.CreateWallet(res.UserId, 1)
			if rerr != nil {
				return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
			}
			wallets = append(wallets, *wallet)
		}
		if len(wallets) == 0 {
			fmt.Println("Wallet User id ", res.UserId)

			wallets, rerr = c.walletService.GetWalletsByUserId(res.UserId)
			if rerr != nil {
				fmt.Println(rerr.Code)
				if rerr.Code == 5 {
					wallet, rerr := c.walletService.CreateWallet(res.UserId, 1)
					if rerr != nil {
						return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
					}
					wallets = append(wallets, *wallet)
				} else {
					return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
				}
			}
		}

		return ctx.JSON(map[string]interface{}{
			"data":    res,
			"profile": profile,
			"wallets": wallets,
			"success": true,
		})
	}
	return ctx.JSON(map[string]interface{}{
		"message": "Code has been sent",
		"success": true,
	})
}

func (c *authenticationController) Signin(ctx *fiber.Ctx) error {
	signinDTO := new(models.SigninDTO)
	if err := ctx.BodyParser(signinDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #101",
			"success": false,
		})
	}

	signinDTO.Ip = ctx.IP()
	signinDTO.Agent = string(ctx.Context().UserAgent())
	res, err := c.authenticationService.Signin(signinDTO)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	if signinDTO.Signin_Method == 0 {
		return ctx.JSON(map[string]interface{}{
			"data":    res,
			"success": true,
		})
	}
	if res != nil {
		profile, rerr := c.profileService.GetProfileByUserId(res.UserId)
		wallets := make([]models.Wallet, 0)
		if rerr != nil {
			if rerr.Code != 5 {
				return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
			}

			username, err := c.GenerateUsername(&models.SignupDTO{Email: signinDTO.Value})
			if err != nil {
				return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
			}

			profile, rerr = c.profileService.AddProfile(res.UserId, username)
			if rerr != nil {
				return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
			}
			wallet, rerr := c.walletService.CreateWallet(res.UserId, 1)
			if rerr != nil {
				return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
			}
			wallets = append(wallets, *wallet)
		}
		if len(wallets) == 0 {
			wallets, rerr = c.walletService.GetWalletsByUserId(res.UserId)
			if rerr != nil {
				if rerr.Code == 5 {
					wallet, rerr := c.walletService.CreateWallet(res.UserId, 1)
					if rerr != nil {
						return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
					}
					wallets = append(wallets, *wallet)
				} else {
					return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
				}
			}
		}
		return ctx.JSON(map[string]interface{}{
			"data":    res,
			"profile": profile,
			"wallets": wallets,
			"success": true,
		})
	}
	return ctx.JSON(map[string]interface{}{
		"message": "code has been sent",
		"success": true,
	})
}

func (c *authenticationController) Verify(ctx *fiber.Ctx) error {
	verifyDTO := new(models.VerifyDTO)
	if err := ctx.BodyParser(verifyDTO); err != nil {
		fmt.Println(err)
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #102",
			"success": false,
		})
	}
	verifyDTO.Ip = ctx.IP()
	verifyDTO.Agent = string(ctx.Context().UserAgent())

	res, err := c.authenticationService.Verify(verifyDTO)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	profile, rerr := c.profileService.GetProfileByUserId(res.Token.UserId)
	wallets := make([]models.Wallet, 0)

	if rerr != nil {
		if rerr.Code != 5 {
			return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
		}

		username, err := c.GenerateUsername(&models.SignupDTO{Email: res.Value})
		if err != nil {
			return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
		}

		profile, rerr = c.profileService.AddProfile(res.Token.UserId, username)
		if rerr != nil {
			return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
		}
		wallet, rerr := c.walletService.CreateWallet(res.Token.UserId, 1)
		if rerr != nil {
			return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
		}
		wallets = append(wallets, *wallet)
	}

	if len(wallets) == 0 {
		wallets, rerr = c.walletService.GetWalletsByUserId(res.Token.UserId)
		if rerr != nil {
			if rerr.Code == 5 {
				wallet, rerr := c.walletService.CreateWallet(res.Token.UserId, 1)
				if rerr != nil {
					return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
				}
				wallets = append(wallets, *wallet)
			} else {
				return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
			}
		}
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"profile": profile,
		"wallets": wallets,
		"success": true,
	})
}

func (c *authenticationController) ForgotPassword(ctx *fiber.Ctx) error {
	emailDTO := new(models.EmailDTO)
	if err := ctx.BodyParser(emailDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #111",
			"success": false,
		})
	}
	err := c.authenticationService.ForgotPassword(emailDTO.Email)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"message": "email sent",
		"success": true,
	})
}

func (c *authenticationController) RenewToken(ctx *fiber.Ctx) error {
	renewTokenDTO := new(models.RenewTokenDTO)
	if err := ctx.BodyParser(renewTokenDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "wrong body, error code #112",
			"success": false,
		})
	}
	renewTokenDTO.Ip = ctx.IP()
	renewTokenDTO.Agent = string(ctx.Context().UserAgent())

	res, err := c.authenticationService.RenewToken(renewTokenDTO)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})

}

func (c *authenticationController) CreateRandomUsername(len int32) (string, *types.Error) {
	randomString, err := idgen.NextNumericString(int(len))
	if err != nil {
		// return ctx.Status(400).JSON(map[string]interface{}{
		// 	"message": "failed to generate username, error code #100-1",
		// 	"success": false,
		// })
		return "", types.NewBadRequestError("failed to generate username, error code #100-1")
	}
	return randomString, nil
}
