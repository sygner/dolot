package services

import (
	"context"
	"dolott_user_gw_http/internal/models"
	"dolott_user_gw_http/internal/types"
	pb "dolott_user_gw_http/proto/api/profile"
)

type (
	ProfileService interface {
		AddProfile(int32, string) (*models.Profile, *types.Error)
		GetProfileByUserId(int32) (*models.Profile, *types.Error)
		CheckUsernameExists(string) (bool, *types.Error)
		GetProfileUsername(string) (*models.Profile, *types.Error)
		UpdateProfile(int32, string) *types.Error
		GetProfileSid(string) (*models.Profile, *types.Error)
	}
	profileService struct {
		profileClient pb.ProfileServiceClient
	}
)

func NewProfileService(profileClient pb.ProfileServiceClient) ProfileService {
	return &profileService{
		profileClient: profileClient,
	}
}

func (c *profileService) AddProfile(userId int32, username string) (*models.Profile, *types.Error) {
	res, err := c.profileClient.AddProfile(context.Background(), &pb.AddProfileRequest{
		UserId:   userId,
		Username: username,
	})

	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}
	return toProfileProto(res)
}

func (c *profileService) GetProfileByUserId(userId int32) (*models.Profile, *types.Error) {
	res, err := c.profileClient.GetProfileByUserId(context.Background(), &pb.UserId{
		UserId: userId,
	})

	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}
	return toProfileProto(res)
}

func (c *profileService) CheckUsernameExists(username string) (bool, *types.Error) {
	_, err := c.profileClient.CheckUsernameExists(context.Background(), &pb.Username{
		Username: username,
	})
	if err != nil {
		rerr := types.ExtractGRPCErrDetails(err)
		if rerr.Code == 5 {
			return false, nil
		} else {
			return false, rerr
		}
	}
	return true, nil
}

func (c *profileService) GetProfileUsername(username string) (*models.Profile, *types.Error) {
	res, err := c.profileClient.GetProfileByUsername(context.Background(), &pb.Username{
		Username: username,
	})

	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}
	return toProfileProto(res)
}

func (c *profileService) GetProfileSid(sid string) (*models.Profile, *types.Error) {
	res, err := c.profileClient.GetProfileBySid(context.Background(), &pb.Sid{
		Sid: sid,
	})

	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}
	return toProfileProto(res)
}

func (c *profileService) ChangeUserScore(userId int32, score float32, increment bool) *types.Error {
	_, err := c.profileClient.ChangeUserScore(context.Background(), &pb.ChangeScoreRequest{
		UserId:    userId,
		Score:     score,
		Increment: increment,
	})

	if err != nil {
		return types.ExtractGRPCErrDetails(err)
	}
	return nil
}

func (c *profileService) ChangeUserGamesQuantity(userId int32, increment bool) *types.Error {
	_, err := c.profileClient.ChangeUserGamesQuantity(context.Background(), &pb.ChangeUserGamesRequest{
		UserId:    userId,
		Increment: increment,
	})

	if err != nil {
		return types.ExtractGRPCErrDetails(err)
	}
	return nil
}

func (c *profileService) ChangeUserWonGames(userId int32, increment bool) *types.Error {
	_, err := c.profileClient.ChangeUserWonGames(context.Background(), &pb.ChangeUserGamesRequest{
		UserId:    userId,
		Increment: increment,
	})

	if err != nil {
		return types.ExtractGRPCErrDetails(err)
	}
	return nil
}

func (c *profileService) ChangeUserLostGames(userId int32, increment bool) *types.Error {
	_, err := c.profileClient.ChangeUserLostGames(context.Background(), &pb.ChangeUserGamesRequest{
		UserId:    userId,
		Increment: increment,
	})

	if err != nil {
		return types.ExtractGRPCErrDetails(err)
	}
	return nil
}

func (c *profileService) ChangeUserRank(userId, rankAmount int32, increment bool) *types.Error {
	_, err := c.profileClient.ChangeUserRank(context.Background(), &pb.ChangeUserRankRequest{
		UserId:     userId,
		RankAmount: rankAmount,
		Increment:  increment,
	})

	if err != nil {
		return types.ExtractGRPCErrDetails(err)
	}
	return nil
}

func (c *profileService) ChangeUserImpression(userId, impression int32, increment bool) *types.Error {
	_, err := c.profileClient.ChangeUserImpression(context.Background(), &pb.ChangeImpressionRequest{
		UserId:     userId,
		Impression: impression,
		Increment:  increment,
	})

	if err != nil {
		return types.ExtractGRPCErrDetails(err)
	}
	return nil
}

func (c *profileService) UpdateProfile(userId int32, username string) *types.Error {
	_, err := c.profileClient.UpdateProfile(context.Background(), &pb.UpdateProfileRequest{
		UserId:   userId,
		Username: username,
	})

	if err != nil {
		return types.ExtractGRPCErrDetails(err)
	}
	return nil
}
