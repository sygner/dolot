package services

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"dolott_authentication/internal/global"
	"dolott_authentication/internal/models"
	"dolott_authentication/internal/pkg"
	"dolott_authentication/internal/repository"
	"dolott_authentication/internal/types"
	"dolott_authentication/internal/utils"

	"safir/libs/idgen"

	"github.com/redis/go-redis/v9"
)

type (
	AuthenticationService interface {
		GetUserByEmail(string) (*models.User, *types.Error)
		GetUserByUserId(int32) (*models.User, *types.Error)
		Signup(*models.UserDTO) (*models.Token, *types.Error)
		Verify(data *models.VerifyDTO) (*models.VerifyResponse, *types.Error)
		Signin(*models.LoginDTO) (*models.Token, *types.Error)
	}

	authenticationService struct {
		repository repository.AuthenticationRepository
		rdDB       *redis.Client
	}
)

func NewAuthenticationService(repository repository.AuthenticationRepository, rdDB *redis.Client) AuthenticationService {
	return &authenticationService{
		repository: repository,
		rdDB:       rdDB,
	}
}

func (s *authenticationService) GetUserByEmail(email string) (*models.User, *types.Error) {
	return s.repository.GetUserByEmail(email)
}

func (s *authenticationService) GetUserByUserId(userID int32) (*models.User, *types.Error) {
	return s.repository.GetUserByUserId(userID)
}

func (s *authenticationService) Signup(data *models.UserDTO) (*models.Token, *types.Error) {
	var exists bool
	if !strings.Contains(data.Email, "@") {
		data.AccountUsername = data.Email
		data.Email = ""
		ex, err := s.repository.UserExistsByAccountUsername(data.AccountUsername)
		if err != nil {
			return nil, err
		}
		exists = ex
	} else {
		ex, err := s.repository.UserExistsByEmail(data.Email)
		if err != nil {
			return nil, err
		}
		exists = ex
	}

	if exists {
		if data.IsSSO {
			res, err := s.Signin(&models.LoginDTO{
				Value:       data.Email,
				Agent:       data.Agent,
				Ip:          data.Ip,
				IsSSO:       data.IsSSO,
				Provider:    data.Provider,
				LoginMethod: 1,
			})
			return res, err
		}

		return nil, types.NewBadRequestError("email already exists, error code #2001")
	}

	if data.IsSSO {
		fmt.Println(data.Email, data.AccountUsername)
		fmt.Println("Signup")
		user, rerr := s.repository.AddUser(data)
		if rerr != nil {
			return nil, rerr
		}

		token, rerr := s.createToken(*data, user.UserId, user.UserRole)
		if rerr != nil {
			return nil, rerr
		}

		rerr = s.repository.AddToken(token)
		if rerr != nil {
			return nil, rerr
		}
		rerr = s.repository.AddLoginHistory(user.UserId)
		if rerr != nil {
			return nil, rerr
		}
		return token, nil
	}

	valid, err := utils.ValidatePassword(*data.Password)
	if err != nil {
		return nil, err
	}
	if !valid {
		return nil, types.NewBadRequestError("Password must be at least 8 characters long, contain at least one uppercase letter, one lowercase letter, and at least one number or symbol, error code #2002")
	}

	ctx := context.Background()
	_, redisErr := s.rdDB.Get(ctx, data.Email).Result()
	if redisErr == redis.Nil {
		randNum, redisErr := idgen.NextNumericint32(72112, 98998)
		if redisErr != nil {
			return nil, types.NewInternalError("internal issue, error code #2005")
		}

		hashedPassword, err := utils.HashPassword([]byte(*data.Password))
		if err != nil {
			return nil, err
		}
		data.Password = &hashedPassword

		userData, redisErr := json.Marshal(data)
		if redisErr != nil {
			return nil, types.NewInternalError("internal issue, error code #2007")
		}

		_, redisErr = s.rdDB.Set(ctx, fmt.Sprintf("%d", randNum), userData, time.Second*120).Result()
		if redisErr != nil {
			fmt.Println(redisErr)
			return nil, types.NewInternalError("internal issue, error code #2006")
		}
		fmt.Println(global.TEMPLATE_FILE_PATH)
		err = pkg.NewEmail(data.Email, "Authorization/Signup").SendAuthEmail(global.TEMPLATE_FILE_PATH, fmt.Sprintf("%d", randNum))
		if err != nil {
			return nil, err
		}
	} else if redisErr != nil {
		return nil, types.NewInternalError("internal issue, error code #2003")
	} else {
		return nil, types.NewBadRequestError("process ongoing, error code #2004")
	}

	return nil, nil
}

func (s *authenticationService) Verify(data *models.VerifyDTO) (*models.VerifyResponse, *types.Error) {
	ctx := context.Background()

	// Handle the different verification methods
	switch data.VerifyMethod {
	case 0:
		return s.verifyNewUser(ctx, data.Code, data.Agent)
	case 1:
		return s.verifyExistingUser(ctx, data.Code, data.Agent)
	}

	// Ensure the new password is provided
	if data.NewPassword == nil || *data.NewPassword == "" {
		return nil, types.NewBadRequestError("password not provided, error code #2042")
	}

	// Fetch verification code from Redis and check if it exists
	res, err := s.rdDB.GetDel(ctx, data.Code).Result()
	if err == redis.Nil {
		return nil, types.NewBadRequestError("code not found, error code #2039")
	} else if err != nil {
		return nil, types.NewInternalError("failed to retrieve code from Redis, error code #2040")
	}

	// Unmarshal the user data from the Redis response
	var user models.User
	if err := json.Unmarshal([]byte(res), &user); err != nil {
		return nil, types.NewInternalError("failed to unmarshal user data, error code #2016")
	}

	// Handle SSO case: If user and DTO both indicate SSO and no password exists, allow password creation
	if user.IsSSO {
		_, rerr := s.repository.GetPasswordByUserId(user.UserId)
		if rerr != nil && rerr.Code == 404 { // No password found
			hashedPassword, rerr := utils.HashPassword([]byte(*data.NewPassword))
			if rerr != nil {
				return nil, types.NewInternalError("failed to hash password, error code #2045")
			}

			rerr = s.repository.AddPassword(&models.UserPassword{
				UserId:   user.UserId,
				Password: hashedPassword,
			})
			if rerr != nil {
				return nil, rerr
			}
		} else if rerr != nil {
			return nil, rerr
		} else {
			return nil, types.NewBadRequestError("password already exists for this user, error code #2044")
		}
	}

	// Hash the new password
	hashedPassword, rerr := utils.HashPassword([]byte(*data.NewPassword))
	if rerr != nil {
		return nil, types.NewInternalError("failed to hash password, error code #2045")
	}

	// Update the user's password
	rerr = s.repository.UpdatePassword(user.UserId, hashedPassword)
	if rerr != nil {
		return nil, rerr
	}

	// Delete old tokens and create new ones
	rerr = s.repository.DeleteUserTokens(user.UserId)
	if rerr != nil {
		return nil, types.NewInternalError("failed to delete user tokens, error code #2046")
	}

	// Create a new token
	token, rerr := s.createToken(models.UserDTO{Agent: data.Agent, Ip: data.Ip}, user.UserId, user.UserRole)
	if rerr != nil {
		return nil, rerr
	}

	// Store the new token
	rerr = s.repository.AddToken(token)
	if rerr != nil {
		return nil, types.NewInternalError("failed to store new token, error code #2047")
	}

	// Add login history
	rerr = s.repository.AddLoginHistory(user.UserId)
	if rerr != nil {
		return nil, types.NewInternalError("failed to update login history, error code #2048")
	}

	return &models.VerifyResponse{Token: *token, Value: user.Email}, nil
}

func (s *authenticationService) verifyNewUser(ctx context.Context, code, agent string) (*models.VerifyResponse, *types.Error) {
	res, err := s.rdDB.GetDel(ctx, code).Result()
	if err == redis.Nil {
		return nil, types.NewBadRequestError("code not found, error code #2008")
	} else if err != nil {
		return nil, types.NewInternalError("internal issue, error code #2009")
	}

	var data models.UserDTO
	err = json.Unmarshal([]byte(res), &data)
	if err != nil {
		return nil, types.NewInternalError("internal issue, error code #2010")
	}
	if data.Agent != agent {
		return nil, types.NewBadRequestError("failed to verify. different agent, error code #2014")
	}

	exists, rerr := s.repository.UserExistsByEmail(data.Email)
	if rerr != nil {
		return nil, rerr
	}
	if exists {
		return nil, types.NewBadRequestError("email already exists, error code #2028")
	}

	user, rerr := s.repository.AddUser(&data)
	if rerr != nil {
		return nil, rerr
	}

	rerr = s.repository.AddPassword(&models.UserPassword{UserId: user.UserId, Password: *data.Password})
	if rerr != nil {
		return nil, rerr
	}

	token, rerr := s.createToken(data, user.UserId, user.UserRole)
	if rerr != nil {
		return nil, rerr
	}

	rerr = s.repository.AddToken(token)
	if rerr != nil {
		return nil, rerr
	}
	rerr = s.repository.AddLoginHistory(user.UserId)
	if rerr != nil {
		return nil, rerr
	}
	return &models.VerifyResponse{Token: *token, Value: user.Email}, nil
}

func (s *authenticationService) verifyExistingUser(ctx context.Context, code, agent string) (*models.VerifyResponse, *types.Error) {
	res, err := s.rdDB.GetDel(ctx, code).Result()
	if err == redis.Nil {
		return nil, types.NewBadRequestError("code not found, error code #2014")
	} else if err != nil {
		return nil, types.NewInternalError("internal issue, error code #2015")
	}

	var data models.UserDTO
	err = json.Unmarshal([]byte(res), &data)
	if err != nil {
		return nil, types.NewInternalError("internal issue, error code #2016")
	}
	if data.Agent != agent {
		return nil, types.NewBadRequestError("failed to verify. different agent, error code #2017")
	}

	exists, rerr := s.repository.UserExistsByUserId(data.UserId)
	if rerr != nil {
		return nil, rerr
	}
	if !exists {
		return nil, types.NewBadRequestError("try to signup, error code #2027")
	}

	token, rerr := s.createToken(data, data.UserId, data.UserRole)
	if rerr != nil {
		return nil, rerr
	}

	rerr = s.repository.AddToken(token)
	if rerr != nil {
		return nil, rerr
	}
	rerr = s.repository.AddLoginHistory(data.UserId)
	if rerr != nil {
		return nil, rerr
	}
	return &models.VerifyResponse{Token: *token, Value: data.Email}, nil
}

func (s *authenticationService) createToken(data models.UserDTO, userID int32, userRole string) (*models.Token, *types.Error) {
	sessionId, err := idgen.NextNumericint32(1102, 32767)
	if err != nil {
		return nil, types.NewInternalError("internal issue, error code #2011")
	}

	accessToken, err := idgen.NextAlphanumericString(40)
	if err != nil {
		return nil, types.NewInternalError("internal issue, error code #2012")
	}
	refreshToken, err := idgen.NextAlphanumericString(40)
	if err != nil {
		return nil, types.NewInternalError("internal issue, error code #2013")
	}

	return &models.Token{
		AccessToken:          accessToken,
		RefreshToken:         refreshToken,
		UserId:               userID,
		UserRole:             userRole,
		SessionId:            sessionId,
		TokenStatus:          models.LIVE.String(),
		Ip:                   data.Ip,
		Agent:                data.Agent,
		CreatedAt:            time.Now(),
		AccessTokenExpireAt:  time.Now().Add(time.Minute * 80),
		RefreshTokenExpireAt: time.Now().AddDate(0, 1, 0),
	}, nil
}

func (s *authenticationService) Signin(data *models.LoginDTO) (*models.Token, *types.Error) {
	if data.LoginMethod == 1 {
		return s.signinByEmail(data)
	}
	return s.signinWithPassword(data)
}

func (s *authenticationService) signinByEmail(data *models.LoginDTO) (*models.Token, *types.Error) {
	if !strings.Contains(data.Value, "@") && !data.IsSSO {
		return nil, types.NewBadRequestError("incorrect email, error code #2024")
	}

	user, err := s.repository.GetUserByEmail(data.Value)
	if err != nil {
		if err.Code == 404 && data.IsSSO {
			token, err := s.Signup(&models.UserDTO{
				Email:      data.Value,
				UserRole:   models.USER.String(),
				UserStatus: models.ONGOING.String(),
				Agent:      data.Agent,
				Provider:   data.Provider,
				IsSSO:      data.IsSSO,
				Ip:         data.Ip,
			})
			return token, err
		}
		return nil, err
	}
	err = utils.ValidateUserStatus(user.UserStatus)
	if err != nil {
		return nil, err
	}
	if user.IsSSO && data.IsSSO {
		if data.Provider == nil {
			return nil, types.NewBadRequestError("provider is empty, error code #2073")
		}
		if user.Provider != strings.ToLower(*data.Provider) {
			return nil, types.NewBadRequestError("incorrect provider, error code #2072")
		}
		token, err := s.createToken(models.UserDTO{Agent: data.Agent, Ip: data.Ip}, user.UserId, user.UserRole)
		if err != nil {
			return nil, err
		}

		err = s.repository.AddToken(token)
		if err != nil {
			return nil, err
		}
		err = s.repository.AddLoginHistory(user.UserId)
		if err != nil {
			return nil, err
		}

		return token, nil
	}

	ctx := context.Background()
	randNum, rerr := idgen.NextNumericint32(42112, 68998)
	if rerr != nil {
		return nil, types.NewInternalError("internal issue, error code #2021")
	}

	userDTO := models.UserDTO{
		Email:      user.Email,
		UserId:     user.UserId,
		UserRole:   user.UserRole,
		UserStatus: user.UserStatus,
		Agent:      data.Agent,
		Ip:         data.Ip,
	}

	userData, rerr := json.Marshal(userDTO)
	if rerr != nil {
		return nil, types.NewInternalError("internal issue, error code #2022")
	}

	_, rerr = s.rdDB.Set(ctx, fmt.Sprintf("%d", randNum), userData, time.Second*120).Result()
	if rerr != nil {
		fmt.Println(rerr)
		return nil, types.NewInternalError("internal issue, error code #2023")
	}
	fmt.Println(global.TEMPLATE_FILE_PATH)

	err = pkg.NewEmail(user.Email, "Authorization/Signin").SendAuthEmail(global.TEMPLATE_FILE_PATH, fmt.Sprintf("%d", randNum))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *authenticationService) signinWithPassword(data *models.LoginDTO) (*models.Token, *types.Error) {
	if !strings.Contains(data.Value, "@") {
		return nil, types.NewBadRequestError("incorrect email, error code #2026")
	}
	user, err := s.repository.GetUserByEmail(data.Value)
	if err != nil {
		return nil, err
	}
	err = utils.ValidateUserStatus(user.UserStatus)
	if err != nil {
		return nil, err
	}

	password, err := s.repository.GetPasswordByUserId(user.UserId)
	if err != nil {
		return nil, err
	}

	valid, err := utils.VerifyPassword([]byte(data.Password), []byte(password.Password))
	if err != nil {
		return nil, err
	}
	if !valid {
		return nil, types.NewBadRequestError("password or email is not valid, error code #2025")
	}

	token, err := s.createToken(models.UserDTO{Agent: data.Agent, Ip: data.Ip}, user.UserId, user.UserRole)
	if err != nil {
		return nil, err
	}

	err = s.repository.AddToken(token)
	if err != nil {
		return nil, err
	}
	err = s.repository.AddLoginHistory(user.UserId)
	if err != nil {
		return nil, err
	}

	return token, nil
}
