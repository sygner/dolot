package services

import (
	"context"
	"dolott_authentication/internal/global"
	"dolott_authentication/internal/models"
	"dolott_authentication/internal/pkg"
	"dolott_authentication/internal/repository"
	"dolott_authentication/internal/types"
	"dolott_authentication/internal/utils"
	"encoding/json"
	"fmt"
	"neo/libs/idgen"
	"time"

	"github.com/redis/go-redis/v9"
)

type (
	UserService interface {
		GetUserByUserId(int32) (*models.User, *types.Error)
		GetUsers(*models.Pagination) (*models.Users, *types.Error)
		ChangeUserStatus(int32, string) *types.Error
		GetUserByEmail(string) (*models.User, *types.Error)
		GetUserByAccountUsername(string) (*models.User, *types.Error)
		ResetPassword(*models.ResetPasswordDTO) *types.Error
		ForgotPassword(string) *types.Error
	}
	userService struct {
		repository repository.AuthenticationRepository
		rdDB       *redis.Client
	}
)

func NewUserService(repository repository.AuthenticationRepository, rdDB *redis.Client) UserService {
	return &userService{
		repository: repository,
		rdDB:       rdDB,
	}
}

func (c *userService) GetUserByUserId(userId int32) (*models.User, *types.Error) {
	return c.repository.GetUserByUserId(userId)
}

func (c *userService) GetUsers(data *models.Pagination) (*models.Users, *types.Error) {
	res, err := c.repository.GetUsers(data)
	if err != nil {
		return nil, err
	}
	var total *int32
	if data.Total {
		count, err := c.repository.GetUserCount()
		if err != nil {
			return nil, err
		}

		total = &count
	}
	return &models.Users{Users: res, Total: total}, nil
}

func (c *userService) ChangeUserStatus(userId int32, newStatus string) *types.Error {
	exists, rerr := c.repository.UserExistsByUserId(userId)
	if rerr != nil {
		return rerr
	}
	if !exists {
		return types.NewBadRequestError("user not found, error code #2029")
	}
	switch newStatus {
	case "OnGoing":
		newStatus = models.ONGOING.String()
	case "Suspended":
		newStatus = models.SUSPENDED.String()
	case "BANNED":
		newStatus = models.BANNED.String()
	default:
		return types.NewBadRequestError("status is not valid, error code #2030")

	}
	err := c.repository.ChangeUserStatus(userId, newStatus)
	if err != nil {
		return err
	}

	err = c.repository.DeleteUserTokens(userId)
	if err != nil {
		return err
	}

	return nil
}

func (c *userService) GetUserByEmail(email string) (*models.User, *types.Error) {
	return c.repository.GetUserByEmail(email)
}
func (c *userService) GetUserByAccountUsername(accountUsername string) (*models.User, *types.Error) {
	return c.repository.GetUserByAccountUsername(accountUsername)
}

func (c *userService) ResetPassword(data *models.ResetPasswordDTO) *types.Error {
	user, err := c.repository.GetUserByUserId(data.UserId)
	if err != nil {
		return err
	}
	err = utils.ValidateUserStatus(user.UserStatus)
	if err != nil {
		return err
	}

	password, err := c.repository.GetPasswordByUserId(user.UserId)
	if err != nil {
		return err
	}

	valid, err := utils.VerifyPassword([]byte(data.CurrentPassword), []byte(password.Password))
	if err != nil {
		return err
	}
	if !valid {
		return types.NewBadRequestError("password is not valid, error code #2034")
	}

	hashedPassword, err := utils.HashPassword([]byte(data.NewPassword))
	if err != nil {
		return err
	}
	err = c.repository.UpdatePassword(data.UserId, hashedPassword)
	if err != nil {
		return err
	}
	return c.repository.DeleteUserTokens(data.UserId)
}

func (c *userService) ForgotPassword(email string) *types.Error {
	user, err := c.repository.GetUserByEmail(email)
	if err != nil {
		return err
	}
	err = utils.ValidateUserStatus(user.UserStatus)
	if err != nil {
		return err
	}
	var emailTitle string
	if user.IsSSO {
		res, err := c.repository.GetPasswordByUserId(user.UserId)
		if err != nil {
			if err.Code == 404 {
				emailTitle = "Authorization/SetPassword"
			} else {
				return err

			}
		}
		if res != nil {
			emailTitle = "Authorization/ForgotPassword"
		}
	}
	userData, redisErr := json.Marshal(user)
	if redisErr != nil {
		return types.NewInternalError("internal issue, error code #2036")
	}
	randNum, redisErr := idgen.NextNumericint32(12112, 38998)
	if redisErr != nil {
		return types.NewInternalError("internal issue, error code #2037")
	}
	ctx := context.Background()

	_, redisErr = c.rdDB.Set(ctx, fmt.Sprintf("%d", randNum), userData, time.Second*120).Result()
	if redisErr != nil {
		fmt.Println(redisErr)
		return types.NewInternalError("internal issue, error code #2038")
	}

	err = pkg.NewEmail(user.Email, emailTitle).SendAuthEmail(global.TEMPLATE_FILE_PATH, fmt.Sprintf("%d", randNum))
	if err != nil {
		return err
	}
	return nil
}
