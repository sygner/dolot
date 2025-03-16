package handlers

import (
	"dolott_game/internal/models"
	"dolott_game/internal/types"
	pb "dolott_game/proto/api"
	"strconv"
	"time"
)

func toGamesProto(res *models.Games) *pb.Games {
	games := make([]*pb.Game, 0)
	for _, game := range res.Games {
		games = append(games, toGameProto(&game))
	}
	return &pb.Games{Games: games, Total: res.Total}
}

func toGameProto(res *models.Game) *pb.Game {
	return &pb.Game{
		NumMainNumbers:   res.NumMainNumbers,
		NumBonusNumbers:  res.NumBonusNumbers,
		MainNumberRange:  res.MainNumberRange,
		BonusNumberRange: res.BonusNumberRange,
		CreatorId:        res.CreatorId,
		Id:               res.Id,
		Name:             res.Name,
		Result:           res.Result,
		GameType:         pb.GameType(models.FromString(res.GameType)),
		Prize:            res.Prize,
		AutoCompute:      res.AutoCompute,
		StartTime:        strconv.FormatInt(res.StartTime.Unix(), 10),
		EndTime:          strconv.FormatInt(res.EndTime.Unix(), 10),
		CreatedAt:        strconv.FormatInt(res.CreatedAt.Unix(), 10),
	}
}

func toUserChoiceResultDetailsProto(res []models.UserChoiceResultDetail) []*pb.UserChoiceResultDetail {
	choices := make([]*pb.UserChoiceResultDetail, 0)
	for _, division := range res {
		choices = append(choices, toUserChoiceResultDetailProto(&division))
	}
	return choices
}
func toUserChoiceResultDetailProto(res *models.UserChoiceResultDetail) *pb.UserChoiceResultDetail {
	return &pb.UserChoiceResultDetail{
		UserId:            res.UserId,
		ChosenMainNumber:  res.ChosenMainNumbers,
		ChosenBonusNumber: res.ChosenBonusNumber,
		MatchCount:        res.MatchCount,
	}
}

func toDivisionResultsProto(res []models.DivisionResult) *pb.DivisionResults {
	divisions := make([]*pb.DivisionResult, 0)
	for _, division := range res {
		divisions = append(divisions, toDivisionResultProto(&division))
	}
	return &pb.DivisionResults{DivisionResults: divisions}
}
func toDivisionResultProto(res *models.DivisionResult) *pb.DivisionResult {
	return &pb.DivisionResult{
		HasBonus:   res.HasBonus,
		UserChoice: toUserChoiceResultDetailsProto(res.UserChoices),
		MatchCount: res.MatchCount,
		Division:   res.Division,
	}
}

func toWinnerResultProto(res *models.Winners) *pb.Winner {
	return &pb.Winner{
		Id:           res.Id,
		GameId:       res.GameId,
		GameType:     res.GameType,
		Divisions:    toDivisionResultsProto(res.Divisions),
		ResultNumber: res.ResultNumber,
		TotalPaid:    res.TotalPaid,
		Prize:        uint32(res.Prize),
		Jackpot:      res.JackPot,
		CreatedAt:    strconv.FormatInt(res.CreatedAt.Unix(), 10),
	}
}

func toUserChoisesProto(res models.UserChoices) *pb.UserChoices {
	users := make([]*pb.UserChoice, 0)
	for _, user := range res.UserChoices {
		users = append(users, toUserChoiceProto(&user))
	}
	return &pb.UserChoices{UserChoices: users, Total: res.Total}
}

func toUserChoiceProto(res *models.UserChoice) *pb.UserChoice {
	outMainNumbers := make([]*pb.ChosenMainNumbers, 0)
	for _, c := range res.ChosenMainNumbers {
		outMainNumbers = append(outMainNumbers, &pb.ChosenMainNumbers{
			ChosenMainNumbers: c,
		})
	}
	outBonusNumbers := make([]*pb.ChosenBonusNumbers, 0)
	for _, c := range res.ChosenBonusNumbers {
		outBonusNumbers = append(outBonusNumbers, &pb.ChosenBonusNumbers{
			ChosenBonusNumbers: c,
		})
	}
	return &pb.UserChoice{
		Id:                 res.Id,
		UserId:             res.UserId,
		GameId:             res.GameId,
		ChosenMainNumbers:  outMainNumbers,
		ChosenBonusNumbers: outBonusNumbers,
		CreatedAt:          res.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
func parseTime(timeStr string, errMsg string) (time.Time, error) {
	timeInt64, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		rerr := types.NewInternalError(errMsg)
		return time.Time{}, rerr.ErrorToGRPCStatus()
	}
	return time.Unix(timeInt64, 0), nil
}
