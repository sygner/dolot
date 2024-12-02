package services

import (
	"dolott_user_gw_http/internal/models"
	"dolott_user_gw_http/internal/types"
	"dolott_user_gw_http/internal/utils"
	game_pb "dolott_user_gw_http/proto/api/game"
	profile_pb "dolott_user_gw_http/proto/api/profile"
)

func toDivisionResultsProto(res []*game_pb.DivisionResult) []models.DivisionResult {
	divisions := make([]models.DivisionResult, 0)
	for _, division := range res {
		divisions = append(divisions, toDivisionResultProto(division))
	}
	return divisions
}
func toDivisionResultProto(res *game_pb.DivisionResult) models.DivisionResult {
	return models.DivisionResult{
		HasBonus:    res.HasBonus,
		UserChoices: toUserChoiceResultDetailsProto(res.UserChoice),
		MatchCount:  res.MatchCount,
	}
}

func toUserChoiceResultDetailsProto(res []*game_pb.UserChoiceResultDetail) []models.UserChoiceResultDetail {
	choices := make([]models.UserChoiceResultDetail, 0)
	for _, division := range res {
		choices = append(choices, toUserChoiceResultDetailProto(division))
	}
	return choices
}
func toUserChoiceResultDetailProto(res *game_pb.UserChoiceResultDetail) models.UserChoiceResultDetail {
	return models.UserChoiceResultDetail{
		UserId:            res.UserId,
		ChosenMainNumbers: res.ChosenMainNumber,
		ChosenBonusNumber: res.ChosenBonusNumber,
		MatchCount:        res.MatchCount,
	}
}

func toGamesProto(res *game_pb.Games) (*models.Games, *types.Error) {
	games := make([]models.Game, 0)
	for _, game := range res.Games {
		g, err := toGameProto(game)
		if err != nil {
			return nil, err
		}
		games = append(games, *g)
	}
	return &models.Games{Games: games, Total: res.Total}, nil
}

func toGameProto(res *game_pb.Game) (*models.Game, *types.Error) {
	createdAt, err := utils.ParseTime(res.CreatedAt, "failed to convert the created at, wrong format #1-1")
	if err != nil {
		return nil, err
	}

	startTime, err := utils.ParseTime(res.StartTime, "failed to convert the start time, wrong format #1-2")
	if err != nil {
		return nil, err
	}

	endTime, err := utils.ParseTime(res.EndTime, "failed to convert the end time, wrong format #1-3")
	if err != nil {
		return nil, err
	}
	return &models.Game{
		NumMainNumbers:   res.NumMainNumbers,
		NumBonusNumbers:  res.NumBonusNumbers,
		MainNumberRange:  res.MainNumberRange,
		BonusNumberRange: res.BonusNumberRange,
		CreatorId:        res.CreatorId,
		Id:               res.Id,
		Name:             res.Name,
		Result:           res.Result,
		GameType:         res.GameType.String(),
		StartTime:        startTime,
		EndTime:          endTime,
		CreatedAt:        createdAt,
	}, nil
}

func toUserChoisesProto(res *game_pb.UserChoices) *models.UserChoices {
	users := make([]models.UserChoice, 0)
	for _, user := range res.UserChoices {
		users = append(users, *toUserChoiceProto(user))
	}
	return &models.UserChoices{UserChoices: users, Total: res.Total}
}

func toUserChoiceProto(res *game_pb.UserChoice) *models.UserChoice {
	outMainNumbers := make([][]int32, 0)
	for _, c := range res.ChosenMainNumbers {
		outMainNumbers = append(outMainNumbers, c.ChosenMainNumbers)
	}
	outBonusNumbers := make([][]int32, 0)
	for _, c := range res.ChosenBonusNumbers {
		outBonusNumbers = append(outBonusNumbers, c.ChosenBonusNumbers)
	}
	return &models.UserChoice{
		Id:                 res.Id,
		UserId:             res.UserId,
		GameId:             res.GameId,
		ChosenMainNumbers:  outMainNumbers,
		ChosenBonusNumbers: outBonusNumbers,
		CreatedAt:          res.CreatedAt,
	}
}

func toProfileProto(res *profile_pb.Profile) (*models.Profile, *types.Error) {
	createdAt, err := utils.ParseTime(res.CreatedAt, "failed to convert the created at, wrong format #1-4")
	if err != nil {
		return nil, err
	}
	return &models.Profile{
		UserId:        res.UserId,
		Score:         res.Score,
		Impression:    res.Impression,
		Rank:          res.Rank,
		GamesQuantity: res.GamesQuantity,
		WonGames:      res.WonGames,
		LostGames:     res.LostGames,
		Sid:           res.Sid,
		Username:      res.Username,
		CreatedAt:     createdAt,
	}, nil
}
