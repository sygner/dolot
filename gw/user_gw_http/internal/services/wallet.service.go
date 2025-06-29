package services

import (
	"context"
	"dolott_user_gw_http/internal/admin"
	"dolott_user_gw_http/internal/models"
	"dolott_user_gw_http/internal/types"
	pb "dolott_user_gw_http/proto/api/wallet"
	"fmt"
)

type (
	WalletService interface {
		GetAllCoins() ([]models.Coin, *types.Error)
		GetCoinById(int32) (*models.Coin, *types.Error)
		GetCoinBySymbol(string) (*models.Coin, *types.Error)
		CreateWallet(int32, int32) (*models.Wallet, *types.Error)
		GetBalanceByCoinIdAndUserId(int32, int32) (float64, *types.Error)
		GetWalletsByUserId(int32) ([]models.Wallet, *types.Error)
		GetWalletBySid(string) (*models.Wallet, *types.Error)
		GetWalletByWalletId(int32) (*models.Wallet, *types.Error)
		GetWalletByAddress(string) (*models.Wallet, *types.Error)
		DeleteWalletByWalletId(int32) *types.Error
		DeleteWalletsByUserId(int32) *types.Error
		GetTransactionByTxId(string) (*models.Transaction, *types.Error)
		GetTransactionsByWalletIdAndUserId(int32, int32) (*models.Transactions, *types.Error)
		AddTransaction(*models.AddTransactionDTO) (*models.Transaction, *types.Error)
		GetTransactionsByUserId(int32) (*models.Transactions, *types.Error)
		GetPreTransactionDetail(*models.AddTransactionDTO) (*models.PreTransactionDetail, *types.Error)
		GetTransactionsByWalletIdAndUserIdAndPagination(int32, int32, *models.Pagination) (*models.Transactions, *types.Error)
		GetTransactionsByUserIdAndPagination(int32, *models.Pagination) (*models.Transactions, *types.Error)
	}
	walletService struct {
		walletClient pb.WalletServiceClient
	}
)

func NewWalletService(walletClient pb.WalletServiceClient) WalletService {
	return &walletService{
		walletClient: walletClient,
	}
}

func (c *walletService) GetAllCoins() ([]models.Coin, *types.Error) {
	res, err := c.walletClient.GetAllCoins(context.Background(), &pb.Empty{})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return toCoinsProto(res), nil
}

func (c *walletService) GetCoinById(coinId int32) (*models.Coin, *types.Error) {
	res, err := c.walletClient.GetCoinById(context.Background(), &pb.CoinId{
		CoinId: coinId,
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return toCoinProto(res), nil
}

func (c *walletService) GetCoinBySymbol(symbol string) (*models.Coin, *types.Error) {
	res, err := c.walletClient.GetCoinBySymbol(context.Background(), &pb.CoinSymbol{
		Symbol: symbol,
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return toCoinProto(res), nil
}

func (c *walletService) CreateWallet(userId int32, coinId int32) (*models.Wallet, *types.Error) {
	res, err := c.walletClient.CreateWallet(context.Background(), &pb.CreateWalletRequest{
		UserId: userId,
		CoinId: coinId,
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}
	wallet, rerr := toWalletProto(res)
	if rerr != nil {
		return nil, rerr
	}
	return wallet, nil
}

func (c *walletService) GetBalanceByCoinIdAndUserId(userId int32, coinId int32) (float64, *types.Error) {
	res, err := c.walletClient.GetBalanceByCoinIdAndUserId(context.Background(), &pb.GetWalletByCoinIdAndUserIdRequest{
		UserId: userId,
		CoinId: coinId,
	})
	if err != nil {
		return 0, types.ExtractGRPCErrDetails(err)
	}

	return float64(res.Balance), nil
}

func (c *walletService) GetWalletsByUserId(userId int32) ([]models.Wallet, *types.Error) {
	res, err := c.walletClient.GetWalletsByUserId(context.Background(), &pb.UserId{
		UserId: userId,
	})
	if err != nil {
		fmt.Println("ERROR", err)
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return toWalletsProto(res)
}

func (c *walletService) GetWalletBySid(sid string) (*models.Wallet, *types.Error) {
	res, err := c.walletClient.GetWalletBySid(context.Background(), &pb.Sid{
		Sid: sid,
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return toWalletProto(res)
}

func (c *walletService) GetWalletByWalletId(walletId int32) (*models.Wallet, *types.Error) {
	res, err := c.walletClient.GetWalletByWalletId(context.Background(), &pb.WalletId{
		Id: walletId,
	})

	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return toWalletProto(res)
}

func (c *walletService) GetWalletByAddress(address string) (*models.Wallet, *types.Error) {
	res, err := c.walletClient.GetWalletByAddress(context.Background(), &pb.Address{
		Address: address,
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return toWalletProto(res)
}

func (c *walletService) DeleteWalletByWalletId(walletId int32) *types.Error {
	_, err := c.walletClient.DeleteWalletByWalletId(context.Background(), &pb.WalletId{
		Id: walletId,
	})
	if err != nil {
		return types.ExtractGRPCErrDetails(err)
	}

	return nil
}

func (c *walletService) DeleteWalletsByUserId(userId int32) *types.Error {
	_, err := c.walletClient.DeleteWalletsByUserId(context.Background(), &pb.UserId{
		UserId: userId,
	})
	if err != nil {
		return types.ExtractGRPCErrDetails(err)
	}

	return nil
}

func (c *walletService) CheckTransactionExists(txId string) (bool, *types.Error) {
	res, err := c.walletClient.CheckTransactionExists(context.Background(), &pb.TransactionId{
		TxId: txId,
	})
	if err != nil {
		return false, types.ExtractGRPCErrDetails(err)
	}

	return res.Result, nil
}

func (c *walletService) GetTransactionByTxId(txId string) (*models.Transaction, *types.Error) {
	res, err := c.walletClient.GetTransactionByTxId(context.Background(), &pb.TransactionId{
		TxId: txId,
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return toTransactionProto(res)
}

func (c *walletService) GetTransactionsByWalletIdAndUserId(walletId int32, userId int32) (*models.Transactions, *types.Error) {
	res, err := c.walletClient.GetTransactionsByWalletId(context.Background(), &pb.GetTransactionsByWalletIdRequest{
		UserId:   userId,
		WalletId: walletId,
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return toTransactionsProto(res)
}

func (c *walletService) GetPreTransactionDetail(addTransactionRequest *models.AddTransactionDTO) (*models.PreTransactionDetail, *types.Error) {
	res, err := c.walletClient.GetPreTransactionDetail(context.Background(), &pb.AddTransactionRequest{
		Amount:           addTransactionRequest.Amount,
		FromWalletId:     addTransactionRequest.FromWalletId,
		FromWalletUserId: addTransactionRequest.FromWalletUserId,
		ToWalletAddress:  addTransactionRequest.ToWalletAddress,
		CoinId:           addTransactionRequest.CoinId,
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}
	return &models.PreTransactionDetail{
		GasLimit:      float64(res.GetGasLimit()),
		Amount:        float64(res.GetAmount()),
		Tax:           admin.TRANSACTION_TAX_PERCENTAGE,
		TaxPercentage: fmt.Sprintf("%d", uint8((admin.TRANSACTION_TAX_PERCENTAGE*100))) + "%",
	}, nil
}

func (c *walletService) AddTransaction(addTransactionRequest *models.AddTransactionDTO) (*models.Transaction, *types.Error) {
	res, err := c.walletClient.AddTransaction(context.Background(), &pb.AddTransactionRequest{
		Amount:           addTransactionRequest.Amount,
		FromWalletId:     addTransactionRequest.FromWalletId,
		FromWalletUserId: addTransactionRequest.FromWalletUserId,
		ToWalletAddress:  addTransactionRequest.ToWalletAddress,
		CoinId:           addTransactionRequest.CoinId,
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return toTransactionProto(res)
}

func (c *walletService) GetTransactionsByUserId(userId int32) (*models.Transactions, *types.Error) {
	res, err := c.walletClient.GetTransactionsByUserId(context.Background(), &pb.UserId{
		UserId: userId,
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return toTransactionsProto(res)
}

func (c *walletService) GetTransactionsByWalletIdAndUserIdAndPagination(walletId, userId int32, pagInation *models.Pagination) (*models.Transactions, *types.Error) {
	res, err := c.walletClient.GetTransactionsByWalletIdAndUserIdAndPagination(context.Background(), &pb.GetTransactionsByWalletIdAndUserIdAndPaginationRequest{
		UserId:   userId,
		WalletId: walletId,
		Pagination: &pb.Pagination{
			Offset: pagInation.Offset,
			Limit:  pagInation.Limit,
			Total:  pagInation.Total,
		},
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}
	return toTransactionsProto(res)
}

func (c *walletService) GetTransactionsByUserIdAndPagination(userId int32, pagInation *models.Pagination) (*models.Transactions, *types.Error) {
	res, err := c.walletClient.GetTransactionsByUserIdAndPagination(context.Background(), &pb.GetTransactionsByUserIdAndPaginationRequest{
		UserId: userId,
		Pagination: &pb.Pagination{
			Offset: pagInation.Offset,
			Limit:  pagInation.Limit,
			Total:  pagInation.Total,
		},
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return toTransactionsProto(res)
}
