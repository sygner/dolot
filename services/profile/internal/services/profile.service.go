package services

import (
	"dolott_profile/internal/models"
	"dolott_profile/internal/repository"
	"dolott_profile/internal/types"
	"safir/libs/idgen"
	"time"
)

type (
	ProfileService interface {
		AddProfile(*models.AddProfileDTO) (*models.Profile, *types.Error)

		GetProfileByUsername(string) (*models.Profile, *types.Error)

		GetProfileBySid(string) (*models.Profile, *types.Error)

		GetProfileByUserId(int32) (*models.Profile, *types.Error)

		ChangeUserScore(int32, float32, bool) *types.Error

		ChangeUserGamesQuantity(int32, bool) *types.Error

		ChangeUserWonGames(int32, bool) *types.Error

		ChangeUserLostGames(int32, bool) *types.Error

		AdjustUserRank(int32, int32, bool) *types.Error

		ChangeUserImpression(int32, int32, bool) *types.Error

		CheckUsernameExists(string) *types.Error

		UpdateProfile(int32, string) *types.Error

		GetAllUserRanking(int32) (*models.Ranking, *types.Error)

		SearchUsername(string) ([]models.Profile, *types.Error)

		GetUserLeaderBoard(int32) ([]models.Profile, *types.Error)

		ChangeImpressionAndDCoin(int32, int32, int32) *types.Error
		// 	IncrementScoreByUserId(int32, float32) *types.Error
		// 	DecrementScoreByUserId(int32, float32) *types.Error

		// 	IncrementUserGamesQuantity(int32) *types.Error
		// 	DecrementUserGamesQuantity(int32) *types.Error

		// 	IncrementUserWonGames(int32) *types.Error
		// 	DecrementUserWonGames(int32) *types.Error

		// 	IncrementUserLostGames(int32) *types.Error
		// 	DecrementUserLostGames(int32) *types.Error
	}
	profileService struct {
		repository repository.ProfileRepository
	}
)

func NewProfileRepository(repository repository.ProfileRepository) ProfileService {
	return &profileService{
		repository: repository,
	}
}

func (c *profileService) AddProfile(data *models.AddProfileDTO) (*models.Profile, *types.Error) {
	res, err := c.repository.GetProfileByUserId(data.UserId)
	if err != nil {
		if err.Code != 404 {
			return nil, err
		}
	}
	if res != nil {
		return nil, types.NewAlreadyExistsError("profile exists for the user #3100")
	}

	res, err = c.repository.GetProfileByUsername(data.Username)
	if err != nil {
		if err.Code != 404 {
			return nil, err
		}
	}
	if res != nil {
		return nil, types.NewAlreadyExistsError("this username already exists #3101")
	}

	sid, rerr := idgen.NextNumericString(30)
	if rerr != nil {
		return nil, types.NewBadRequestError("failed to make the sid #3102")
	}

	highestRank, err := c.repository.GetHighestRank()
	if err != nil {
		return nil, err
	}

	data.Sid = sid
	data.Impression = 0
	data.Rank = highestRank + 1
	data.Score = 0
	data.GamesQuantity = 0
	data.LostGames = 0
	data.WonGames = 0

	// err = c.repository.ChangeAllRanks(true)
	// if err != nil {
	// 	return nil, err
	// }
	err = c.repository.AddProfile(data)
	if err != nil {
		err = c.repository.ChangeAllRanks(false)
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	return &models.Profile{
		UserId:        data.UserId,
		Sid:           data.Sid,
		Username:      data.Username,
		Score:         data.Score,
		Impression:    data.Impression,
		Rank:          data.Rank,
		GamesQuantity: data.GamesQuantity,
		CreatedAt:     time.Now(),
	}, nil

}

func (c *profileService) GetProfileByUsername(username string) (*models.Profile, *types.Error) {
	return c.repository.GetProfileByUsername(username)
}

func (c *profileService) GetProfileBySid(sid string) (*models.Profile, *types.Error) {
	return c.repository.GetProfileBySid(sid)
}

func (c *profileService) GetProfileByUserId(userId int32) (*models.Profile, *types.Error) {
	return c.repository.GetProfileByUserId(userId)
}

func (c *profileService) ChangeUserScore(userId int32, score float32, increment bool) *types.Error {
	return c.repository.ChangeUserScore(userId, score, increment)
}

func (c *profileService) ChangeUserGamesQuantity(userId int32, increment bool) *types.Error {
	return c.repository.ChangeUserGamesQuantity(userId, increment)
}

func (c *profileService) ChangeUserWonGames(userId int32, increment bool) *types.Error {
	return c.repository.ChangeUserWonGames(userId, increment)
}

func (c *profileService) ChangeUserLostGames(userId int32, increment bool) *types.Error {
	return c.repository.ChangeUserLostGames(userId, increment)
}

func (c *profileService) AdjustUserRank(userId int32, rankAmount int32, increment bool) *types.Error {
	return c.repository.AdjustUserRank(userId, rankAmount, increment)
}

func (c *profileService) ChangeUserImpression(userId int32, impression int32, increment bool) *types.Error {
	return c.repository.ChangeUserImpression(userId, impression, increment)
}

func (c *profileService) CheckUsernameExists(username string) *types.Error {
	return c.repository.CheckUsernameExists(username)
}

func (c *profileService) UpdateProfile(userId int32, username string) *types.Error {
	res, err := c.repository.GetProfileByUsername(username)
	if err != nil {
		if err.Code != 404 {
			return err
		}
	}
	if res != nil {
		return types.NewAlreadyExistsError("this username already exists #3101")
	}
	return c.repository.UpdateProfile(userId, username)
}

func (c *profileService) GetAllUserRanking(userId int32) (*models.Ranking, *types.Error) {
	return c.repository.GetAllUserRanking(userId)
}

func (c *profileService) SearchUsername(username string) ([]models.Profile, *types.Error) {
	return c.repository.SearchUsername(username)
}

func (c *profileService) GetUserLeaderBoard(userId int32) ([]models.Profile, *types.Error) {
	return c.repository.GetUserLeaderBoard(userId)
}

func (c *profileService) ChangeImpressionAndDCoin(userId int32, impression int32, dCredit int32) *types.Error {
	return c.repository.ChangeImpressionAndDCoin(userId, impression, dCredit)
}

// func (c *profileService) IncrementScoreByUserId(userId int32, score float32) *types.Error {
// 	return c.repository.IncrementUserScore(userId, score)
// }

// func (c *profileService) DecrementScoreByUserId(userId int32, score float32) *types.Error {
// 	return c.repository.DecrementUserScore(userId, score)
// }

// func (c *profileService) IncrementUserGamesQuantity(userId int32) *types.Error {
// 	return c.repository.IncrementUserGamesQuantity(userId)
// }

// func (c *profileService) DecrementUserGamesQuantity(userId int32) *types.Error {
// 	return c.repository.DecrementUserGamesQuantity(userId)
// }

// func (c *profileService) IncrementUserWonGames(userId int32) *types.Error {
// 	return c.repository.IncrementUserWonGames(userId)
// }

// func (c *profileService) DecrementUserWonGames(userId int32) *types.Error {
// 	return c.repository.DecrementUserWonGames(userId)
// }

// func (c *profileService) IncrementUserLostGames(userId int32) *types.Error {
// 	return c.repository.IncrementUserLostGames(userId)
// }

// func (c *profileService) DecrementUserLostGames(userId int32) *types.Error {
// 	return c.repository.DecrementUserLostGames(userId)
// }
