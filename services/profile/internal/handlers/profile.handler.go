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
		DCoin:         res.DCoin,
		Rank:          res.Rank,
		GamesQuantity: res.GamesQuantity,
		Sid:           res.Sid,
		Username:      res.Username,
		CreatedAt:     fmt.Sprintf("%d", res.CreatedAt.Unix()),
		HighestRank:   res.HighestRank,
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
		DCoin:         res.DCoin,
		Rank:          res.Rank,
		GamesQuantity: res.GamesQuantity,
		WonGames:      res.WonGames,
		LostGames:     res.LostGames,
		Sid:           res.Sid,
		Username:      res.Username,
		CreatedAt:     fmt.Sprintf("%d", res.CreatedAt.Unix()),
		HighestRank:   res.HighestRank,
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
		DCoin:         res.DCoin,
		Rank:          res.Rank,
		GamesQuantity: res.GamesQuantity,
		WonGames:      res.WonGames,
		LostGames:     res.LostGames,
		Sid:           res.Sid,
		Username:      res.Username,
		CreatedAt:     fmt.Sprintf("%d", res.CreatedAt.Unix()),
		HighestRank:   res.HighestRank,
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
		DCoin:         res.DCoin,
		Rank:          res.Rank,
		GamesQuantity: res.GamesQuantity,
		WonGames:      res.WonGames,
		LostGames:     res.LostGames,
		Sid:           res.Sid,
		Username:      res.Username,
		CreatedAt:     fmt.Sprintf("%d", res.CreatedAt.Unix()),
		HighestRank:   res.HighestRank,
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

func (c *ProfileHandler) GetAllUserRanking(ctx context.Context, request *pb.UserId) (*pb.Ranking, error) {
	res, err := c.profileService.GetAllUserRanking(request.UserId)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return &pb.Ranking{
		TotalRanking:       res.TotalRanking,
		IndividualRanking:  res.IndividualRanking,
		SeasonRanking:      res.SeasonRanking,
		MonthRanking:       res.MonthRanking,
		SeasonRankingCount: res.SeasonRankChangesCount,
		MonthRankingCount:  res.MonthRankChangesCount,
		TotalRankingCount:  res.AllRankChangesCount,
	}, nil
}

func (c *ProfileHandler) SearchUsername(ctx context.Context, request *pb.Username) (*pb.Profiles, error) {
	res, err := c.profileService.SearchUsername(request.Username)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	var profiles []*pb.Profile
	for _, profile := range res {
		profiles = append(profiles, &pb.Profile{
			UserId:        profile.UserId,
			Score:         profile.Score,
			Impression:    profile.Impression,
			DCoin:         profile.DCoin,
			Rank:          profile.Rank,
			GamesQuantity: profile.GamesQuantity,
			WonGames:      profile.WonGames,
			LostGames:     profile.LostGames,
			Sid:           profile.Sid,
			Username:      profile.Username,
			CreatedAt:     fmt.Sprintf("%d", profile.CreatedAt.Unix()),
			HighestRank:   profile.HighestRank,
		})
	}

	return &pb.Profiles{Profiles: profiles}, nil
}

func (c *ProfileHandler) GetUserLeaderBoard(ctx context.Context, request *pb.UserId) (*pb.Profiles, error) {
	res, err := c.profileService.GetUserLeaderBoard(request.UserId)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	var profiles []*pb.Profile
	for _, profile := range res {
		profiles = append(profiles, &pb.Profile{
			UserId:        profile.UserId,
			Score:         profile.Score,
			Impression:    profile.Impression,
			DCoin:         profile.DCoin,
			Rank:          profile.Rank,
			GamesQuantity: profile.GamesQuantity,
			WonGames:      profile.WonGames,
			LostGames:     profile.LostGames,
			Sid:           profile.Sid,
			Username:      profile.Username,
			CreatedAt:     fmt.Sprintf("%d", profile.CreatedAt.Unix()),
			HighestRank:   profile.HighestRank,
		})
	}

	return &pb.Profiles{Profiles: profiles}, nil
}

func (c *ProfileHandler) ChangeImpressionAndDCoin(ctx context.Context, request *pb.ChangeImpressionAndDCreditRequest) (*pb.Empty, error) {
	err := c.profileService.ChangeImpressionAndDCoin(request.UserId, request.Impression, request.DCoin)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return &pb.Empty{}, nil
}
