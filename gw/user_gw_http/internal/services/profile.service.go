package services

import (
	"context"
	"dolott_user_gw_http/internal/admin"
	"dolott_user_gw_http/internal/constants"
	"dolott_user_gw_http/internal/models"
	"dolott_user_gw_http/internal/types"
	pb "dolott_user_gw_http/proto/api/profile"
	wallet_pb "dolott_user_gw_http/proto/api/wallet"
)

type (
	ProfileService interface {
		AddProfile(int32, string) (*models.Profile, *types.Error)
		GetProfileByUserId(int32) (*models.Profile, *types.Error)
		CheckUsernameExists(string) (bool, *types.Error)
		GetProfileUsername(string) (*models.Profile, *types.Error)
		UpdateProfile(int32, string) *types.Error
		GetProfileSid(string) (*models.Profile, *types.Error)
		ChangeUserRank(int32, int32, bool) *types.Error
		ChangeUserImpression(int32, int32, bool) *types.Error
		ImpressionExchange(int32, uint32, bool) (*models.Transaction, *types.Error)
		SearchUsername(string) ([]models.Profile, *types.Error)
		GetAllUserRanking(int32) (*models.UserProfileRanking, *types.Error)
		GetUserLeaderBoard(int32) ([]models.Profile, *types.Error)
		ChangeImpressionAndDCoin(int32, int32, int32, bool) *types.Error
	}
	profileService struct {
		profileClient pb.ProfileServiceClient
		walletClient  wallet_pb.WalletServiceClient
	}
)

func NewProfileService(profileClient pb.ProfileServiceClient, walletClient wallet_pb.WalletServiceClient) ProfileService {
	return &profileService{
		profileClient: profileClient,
		walletClient:  walletClient,
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

func (c *profileService) ImpressionExchange(userId int32, impressionAmount uint32, toCoin bool) (*models.Transaction, *types.Error) {
	realAmount := impressionAmount / admin.IMPRESSION_EXCHANGE_RATE

	userProfile, rerr := c.GetProfileByUserId(userId)
	if rerr != nil {
		return nil, rerr
	}
	if toCoin && userProfile.Impression < int32(impressionAmount) {
		return nil, types.NewBadRequestError("you don't have enough impression")
	}

	wallets, err := c.walletClient.GetWalletsByUserId(context.Background(), &wallet_pb.UserId{UserId: userId})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	var luncWallet *models.Wallet
	for _, w := range wallets.Wallets {
		if w.CoinId == 1 {
			wallet, rerr := toWalletProto(w)
			if rerr != nil {
				return nil, rerr
			}
			luncWallet = wallet
			break
		}
	}
	if luncWallet == nil {
		return nil, types.NewBadRequestError("you don't have LUNC wallet")
	}

	// if err := c.ChangeUserImpression(userId, int32(impressionAmount), !toCoin); err != nil {
	// 	return nil, err
	// }
	if _, err := c.profileClient.ChangeImpressionAndDCoin(context.Background(), &pb.ChangeImpressionAndDCreditRequest{
		UserId:     userId,
		Impression: int32(impressionAmount),
		DCoin:      0,
	}); err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	fromWalletId, fromWalletUserId, toWalletAddress := constants.MAIN_LUNC_WALLET_ID, constants.MAIN_LUNC_USER_WALLET_ID, luncWallet.Address
	if !toCoin {
		fromWalletId, fromWalletUserId, toWalletAddress = luncWallet.Id, luncWallet.UserId, constants.MAIN_LUNC_WALLET_ADDRESS
	}

	transaction, err := c.walletClient.AddTransaction(context.Background(), &wallet_pb.AddTransactionRequest{
		Amount:           float64(realAmount),
		FromWalletId:     fromWalletId,
		FromWalletUserId: fromWalletUserId,
		ToWalletAddress:  toWalletAddress,
		CoinId:           1,
	})
	if err != nil {
		c.ChangeUserImpression(userId, int32(impressionAmount), toCoin)
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return toTransactionProto(transaction)
}

func (c *profileService) SearchUsername(username string) ([]models.Profile, *types.Error) {
	res, err := c.profileClient.SearchUsername(context.Background(), &pb.Username{
		Username: username,
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	profiles := make([]models.Profile, 0)
	for _, p := range res.Profiles {
		profile, rerr := toProfileProto(p)
		if rerr != nil {
			return nil, rerr
		}
		profiles = append(profiles, *profile)
	}
	return profiles, nil
}

func (c *profileService) GetAllUserRanking(userId int32) (*models.UserProfileRanking, *types.Error) {
	res, err := c.profileClient.GetAllUserRanking(context.Background(), &pb.UserId{
		UserId: userId,
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return &models.UserProfileRanking{
		TotalRanking:           res.TotalRanking,
		IndividualRanking:      res.IndividualRanking,
		SeasonRanking:          res.SeasonRanking,
		MonthRanking:           res.MonthRanking,
		SeasonRankChangesCount: res.SeasonRankingCount,
		MonthRankChangesCount:  res.MonthRankingCount,
		AllRankChangesCount:    res.TotalRankingCount,
	}, nil
}

func (c *profileService) GetUserLeaderBoard(userId int32) ([]models.Profile, *types.Error) {
	res, err := c.profileClient.GetUserLeaderBoard(context.Background(), &pb.UserId{
		UserId: userId,
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	profiles := make([]models.Profile, 0)
	for _, p := range res.Profiles {
		profile, rerr := toProfileProto(p)
		if rerr != nil {
			return nil, rerr
		}
		profiles = append(profiles, *profile)
	}
	return profiles, nil
}

func (c *profileService) ChangeImpressionAndDCoin(userId int32, impression, dCredit int32, increment bool) *types.Error {
	_, err := c.profileClient.ChangeImpressionAndDCoin(context.Background(), &pb.ChangeImpressionAndDCreditRequest{
		UserId:     userId,
		Impression: impression,
		DCoin:      dCredit,
	})

	if err != nil {
		return types.ExtractGRPCErrDetails(err)
	}
	return nil
}
