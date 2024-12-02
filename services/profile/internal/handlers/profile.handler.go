package handlers

import (
	"context"
	"dolott_profile/internal/models"
	"dolott_profile/internal/services"
	pb "dolott_profile/proto/api"
	"fmt"
	"strings"
)

type ProfileHandler struct {
	pb.UnimplementedProfileServiceServer
	profileService services.ProfileService
}

func NewProfileHandler(profileService services.ProfileService) *ProfileHandler {
	return &ProfileHandler{
		profileService: profileService,
	}
}

func (c *ProfileHandler) AddProfile(ctx context.Context, request *pb.AddProfileRequest) (*pb.Profile, error) {
	data := models.AddProfileDTO{
		UserId:   request.UserId,
		Username: strings.ToLower(request.Username),
	}

	res, err := c.profileService.AddProfile(&data)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	return &pb.Profile{
		UserId:        res.UserId,
		Score:         res.Score,
		Impression:    res.Impression,
		Rank:          res.Rank,
		GamesQuantity: res.GamesQuantity,
		Sid:           res.Sid,
		Username:      res.Username,
		CreatedAt:     fmt.Sprintf("%d", res.CreatedAt.Unix()),
	}, nil
}

func (c *ProfileHandler) GetProfileByUsername(ctx context.Context, request *pb.Username) (*pb.Profile, error) {
	res, err := c.profileService.GetProfileByUsername(request.Username)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return &pb.Profile{
		UserId:        res.UserId,
		Score:         res.Score,
		Impression:    res.Impression,
		Rank:          res.Rank,
		GamesQuantity: res.GamesQuantity,
		WonGames:      res.WonGames,
		LostGames:     res.LostGames,
		Sid:           res.Sid,
		Username:      res.Username,
		CreatedAt:     fmt.Sprintf("%d", res.CreatedAt.Unix()),
	}, nil
}

func (c *ProfileHandler) GetProfileBySid(ctx context.Context, request *pb.Sid) (*pb.Profile, error) {
	res, err := c.profileService.GetProfileBySid(request.Sid)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return &pb.Profile{
		UserId:        res.UserId,
		Score:         res.Score,
		Impression:    res.Impression,
		Rank:          res.Rank,
		GamesQuantity: res.GamesQuantity,
		WonGames:      res.WonGames,
		LostGames:     res.LostGames,
		Sid:           res.Sid,
		Username:      res.Username,
		CreatedAt:     fmt.Sprintf("%d", res.CreatedAt.Unix()),
	}, nil
}
func (c *ProfileHandler) GetProfileByUserId(ctx context.Context, request *pb.UserId) (*pb.Profile, error) {
	res, err := c.profileService.GetProfileByUserId(request.UserId)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return &pb.Profile{
		UserId:        res.UserId,
		Score:         res.Score,
		Impression:    res.Impression,
		Rank:          res.Rank,
		GamesQuantity: res.GamesQuantity,
		WonGames:      res.WonGames,
		LostGames:     res.LostGames,
		Sid:           res.Sid,
		Username:      res.Username,
		CreatedAt:     fmt.Sprintf("%d", res.CreatedAt.Unix()),
	}, nil
}

func (c *ProfileHandler) ChangeUserScore(ctx context.Context, request *pb.ChangeScoreRequest) (*pb.Empty, error) {
	err := c.profileService.ChangeUserScore(request.UserId, request.Score, request.Increment)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return &pb.Empty{}, nil
}

func (c *ProfileHandler) ChangeUserGamesQuantity(ctx context.Context, request *pb.ChangeUserGamesRequest) (*pb.Empty, error) {
	err := c.profileService.ChangeUserGamesQuantity(request.UserId, request.Increment)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return &pb.Empty{}, nil
}

func (c *ProfileHandler) ChangeUserWonGames(ctx context.Context, request *pb.ChangeUserGamesRequest) (*pb.Empty, error) {
	err := c.profileService.ChangeUserWonGames(request.UserId, request.Increment)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return &pb.Empty{}, nil
}

func (c *ProfileHandler) ChangeUserLostGames(ctx context.Context, request *pb.ChangeUserGamesRequest) (*pb.Empty, error) {
	err := c.profileService.ChangeUserLostGames(request.UserId, request.Increment)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return &pb.Empty{}, nil
}

func (c *ProfileHandler) ChangeUserRank(ctx context.Context, request *pb.ChangeUserRankRequest) (*pb.Empty, error) {
	err := c.profileService.AdjustUserRank(request.UserId, request.RankAmount, request.Increment)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return &pb.Empty{}, nil
}

func (c *ProfileHandler) ChangeUserImpression(ctx context.Context, request *pb.ChangeImpressionRequest) (*pb.Empty, error) {
	err := c.profileService.ChangeUserImpression(request.UserId, request.Impression, request.Increment)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return &pb.Empty{}, nil
}

func (c *ProfileHandler) CheckUsernameExists(ctx context.Context, request *pb.Username) (*pb.Empty, error) {
	err := c.profileService.CheckUsernameExists(request.Username)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return &pb.Empty{}, nil
}

func (c *ProfileHandler) UpdateProfile(ctx context.Context, request *pb.UpdateProfileRequest) (*pb.Empty, error) {
	err := c.profileService.UpdateProfile(request.UserId, request.Username)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return &pb.Empty{}, nil
}
