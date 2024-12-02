// Package middleware provides middleware functionalities for gRPC services.
package middleware

import (
	"context"
	"dolott_user_gw_http/internal/models"
	"dolott_user_gw_http/internal/types"
	"dolott_user_gw_http/internal/utils"
	pb_token "dolott_user_gw_http/proto/api/authentication"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// MiddlewareService defines the interface for middleware functionalities.
type MiddlewareService interface {
	VerificationMiddleware(ctx *fiber.Ctx) error
}

// middlewareService is an implementation of the MiddlewareService interface.
type middlewareService struct {
	tokenClient pb_token.TokenServiceClient
}

// NewMiddlewareService creates a new MiddlewareService with the provided TokenServiceClient.
func NewMiddlewareService(tokenClient pb_token.TokenServiceClient) MiddlewareService {
	return &middlewareService{
		tokenClient: tokenClient,
	}
}

// VerificationMiddleware is a middleware function that verifies the user's access token using the TokenService.
// It extracts the token from the Authorization header, extracts user agent information, and makes a gRPC call to verify the token.
// If the token is invalid or the verification fails, it aborts the request with an error response.
// If successful, it sets the user ID in the fiber context and allows the request to proceed.
func (c *middlewareService) VerificationMiddleware(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")

	// Check if the token is empty or does not start with "Bearer "
	if token == "" || token[:7] != "Bearer " {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized, token not found in header, error code #999",
			"success": false,
		})
	}

	// Extract the token string (remove "Bearer " prefix)
	token = token[7:]
	// Verify the extracted token with the TokenService
	tokenResult, cerr := c.tokenClient.VerifyToken(context.TODO(), &pb_token.VerifyTokenRequest{
		AccessToken: token,
		Agent:       string(ctx.Context().UserAgent()),
	})
	if cerr != nil {
		err := types.ExtractGRPCErrDetails(cerr)

		return ctx.Status(err.ErrorToHttpStatus()).JSON(map[string]interface{}{
			"error":   err.Message,
			"success": false,
		})
	}
	createdAt, err := utils.GetTimebyTimestamp(tokenResult.CreatedAt)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Message,
			"success": false,
		})
	}
	accessTokenExpireAt, err := utils.GetTimebyTimestamp(tokenResult.AccessTokenExpireAt)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Message,
			"success": false,
		})
	}
	refreshTokenExpireAt, err := utils.GetTimebyTimestamp(tokenResult.RefreshTokenExpireAt)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Message,
			"success": false,
		})
	}
	tokenModel := models.Token{
		AccessToken:          tokenResult.AccessToken,
		RefreshToken:         tokenResult.RefreshToken,
		TokenStatus:          tokenResult.TokenStatus,
		UserRole:             tokenResult.UserRole,
		Ip:                   tokenResult.Ip,
		Agent:                tokenResult.Agent,
		UserId:               tokenResult.UserId,
		SessionId:            tokenResult.SessionId,
		CreatedAt:            *createdAt,
		AccessTokenExpireAt:  *accessTokenExpireAt,
		RefreshTokenExpireAt: *refreshTokenExpireAt,
	}

	// Set the user ID in the Gin context
	ctx.Locals("user_data", tokenModel)
	return ctx.Next()
}
