package services

import (
	"context"
	"dolott_user_gw_http/internal/admin"
	"dolott_user_gw_http/internal/constants"
	"dolott_user_gw_http/internal/models"
	"dolott_user_gw_http/internal/types"
	pb "dolott_user_gw_http/proto/api/ticket"
	wallet_pb "dolott_user_gw_http/proto/api/wallet"
	"fmt"
	"math"
)

type (
	TicketService interface {
		AddTicket(*models.AddTicketDTO) (*models.Ticket, *types.Error)
		GetTicketBySignatureAndUserId(int32, string) (*models.Ticket, *types.Error)
		GetTicketByUserIdAndTicketId(int32, int32) (*models.Ticket, *types.Error)
		GetAllUserTickets(int32, *models.Pagination) (*models.Tickets, *types.Error)
		GetUserOpenTickets(int32, *models.Pagination) (*models.Tickets, *types.Error)
		GetAllUsedTickets(int32, *models.Pagination) (*models.Tickets, *types.Error)
		UseTickets(int32, int32, string) (*models.Tickets, *types.Error)
		AddTickets([]models.AddTicketDTO, bool) (*models.Tickets, *types.Error)
		GetAllUserTicketsByGameId(int32, string) (*models.Tickets, *types.Error)
		GetAllTicketsByGameId(string) (*models.Tickets, *types.Error)
		BuyTickets(int32, int32, int32, bool) (*models.Tickets, *types.Error)
	}
	ticketService struct {
		ticketClient pb.TicketServiceClient
		walletClient wallet_pb.WalletServiceClient
	}
)

func NewTicketService(ticketClient pb.TicketServiceClient, walletClient wallet_pb.WalletServiceClient) TicketService {
	return &ticketService{
		ticketClient: ticketClient,
		walletClient: walletClient,
	}
}

func (c *ticketService) AddTicket(data *models.AddTicketDTO) (*models.Ticket, *types.Error) {
	res, err := c.ticketClient.AddTicket(context.Background(), &pb.AddTicketRequest{
		UserId:     data.UserId,
		TicketType: data.TicketType,
	})

	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return toTicketProto(res)
}

func (c *ticketService) GetTicketBySignatureAndUserId(userId int32, signature string) (*models.Ticket, *types.Error) {
	res, err := c.ticketClient.GetTicketBySignatureAndUserId(context.Background(), &pb.SignatureAndUserId{
		Signature: signature,
		UserId:    userId,
	})

	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return toTicketProto(res)
}

func (c *ticketService) GetTicketByUserIdAndTicketId(userId int32, ticketId int32) (*models.Ticket, *types.Error) {
	res, err := c.ticketClient.GetTicketByUserIdAndTicketId(context.Background(), &pb.TicketIdAndUserId{
		TicketId: ticketId,
		UserId:   userId,
	})

	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return toTicketProto(res)
}

func (c *ticketService) GetAllUserTickets(userId int32, pagInation *models.Pagination) (*models.Tickets, *types.Error) {
	res, err := c.ticketClient.GetAllUserTickets(context.Background(), &pb.UserIdAndPagination{
		UserId: userId,
		Pagination: &pb.Pagination{
			Offset:   pagInation.Offset,
			Limit:    pagInation.Limit,
			GetTotal: pagInation.Total,
		},
	})

	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return toTicketsProtos(res)
}

func (c *ticketService) GetUserOpenTickets(userId int32, pagInation *models.Pagination) (*models.Tickets, *types.Error) {
	res, err := c.ticketClient.GetUserOpenTickets(context.Background(), &pb.UserIdAndPagination{
		UserId: userId,
		Pagination: &pb.Pagination{
			Offset:   pagInation.Offset,
			Limit:    pagInation.Limit,
			GetTotal: pagInation.Total,
		},
	})

	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return toTicketsProtos(res)
}

func (c *ticketService) GetAllUsedTickets(userId int32, pagInation *models.Pagination) (*models.Tickets, *types.Error) {
	res, err := c.ticketClient.GetAllUsedTickets(context.Background(), &pb.UserIdAndPagination{
		UserId: userId,
		Pagination: &pb.Pagination{
			Offset:   pagInation.Offset,
			Limit:    pagInation.Limit,
			GetTotal: pagInation.Total,
		},
	})

	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return toTicketsProtos(res)
}

func (c *ticketService) UseTickets(userId int32, amount int32, gameId string) (*models.Tickets, *types.Error) {
	res, err := c.ticketClient.UseTickets(context.Background(), &pb.UseTicketsRequest{
		UserId:            userId,
		TotalUsingTickets: amount,
		GameId:            gameId,
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return toTicketsProtos(res)
}

func (c *ticketService) AddTickets(data []models.AddTicketDTO, shouldReturn bool) (*models.Tickets, *types.Error) {
	tickets := make([]*pb.AddTicketRequest, 0)
	for _, ticket := range data {
		tickets = append(tickets, &pb.AddTicketRequest{UserId: ticket.UserId, TicketType: ticket.TicketType})
	}
	res, err := c.ticketClient.AddTickets(context.Background(), &pb.AddTicketsRequest{Tickets: tickets, ShouldReturn: shouldReturn})

	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return toTicketsProtos(res)
}

func (c *ticketService) GetAllUserTicketsByGameId(userId int32, gameId string) (*models.Tickets, *types.Error) {
	res, err := c.ticketClient.GetAllUserTicketsByGameId(context.Background(), &pb.GetAllUserTicketsByGameIdRequest{
		GameId: gameId,
		UserId: userId,
	})

	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return toTicketsProtos(res)
}

func (c *ticketService) GetAllTicketsByGameId(gameId string) (*models.Tickets, *types.Error) {
	res, err := c.ticketClient.GetAllTicketsByGameId(context.Background(), &pb.GameId{GameId: gameId})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}

	return toTicketsProtos(res)
}

func (c *ticketService) BuyTickets(userId int32, totalTickets int32, fromWalletId int32, shouldReturn bool) (*models.Tickets, *types.Error) {
	total := math.Round(float64(totalTickets) * admin.TICKET_BUY_RATE)
	if total <= 0 {
		return nil, types.NewBadRequestError("total tickets must be greater than 0")
	}

	balance, err := c.walletClient.GetBalanceByCoinIdAndUserId(context.Background(), &wallet_pb.GetWalletByCoinIdAndUserIdRequest{
		UserId: userId,
		CoinId: 1,
	})
	if err != nil {
		return nil, types.ExtractGRPCErrDetails(err)
	}
	if math.Round(float64(balance.Balance)) < total {
		return nil, types.NewBadRequestError("Not enough balance")
	}
	// JOB_QUEUE.Enqueue(Job{
	// 	Priority: Force,
	// 	Type:     TransactionJob,
	// 	Transaction: models.AddTransactionDTO{
	// 		Amount:           total,
	// 		FromWalletId:     fromWalletId,
	// 		FromWalletUserId: userId,
	// 		ToWalletAddress:  constants.MAIN_WALLET,
	// 		CoinId:           1,
	// 	},
	// })
	_, rerr := c.walletClient.AddTransaction(context.Background(), &wallet_pb.AddTransactionRequest{
		Amount:           total,
		FromWalletId:     fromWalletId,
		FromWalletUserId: userId,
		ToWalletAddress:  constants.MAIN_LUNC_WALLET_ADDRESS,
		CoinId:           1,
	})
	if rerr != nil {
		fmt.Println("W1")
		return nil, types.ExtractGRPCErrDetails(rerr)
	}
	// resTransaction, err := c.walletClient.AddTransaction(context.Background(), &wallet_pb.AddTransactionRequest{
	// 	Amount:           total,
	// 	FromWalletId:     fromWalletId,
	// 	FromWalletUserId: userId,
	// 	CoinId:           1,
	// 	ToWalletAddress:  constants.MAIN_WALLET,
	// })
	// if err != nil {
	// 	return nil, types.ExtractGRPCErrDetails(err)
	// }
	// fmt.Println(resTransaction)
	tickets := make([]*pb.AddTicketRequest, 0)
	for i := 0; i < int(total); i++ {
		tickets = append(tickets, &pb.AddTicketRequest{UserId: userId, TicketType: "purchased"})
	}
	res, err := c.ticketClient.AddTickets(context.Background(), &pb.AddTicketsRequest{
		Tickets:      tickets,
		ShouldReturn: shouldReturn,
	})
	if err != nil {
		userWallets, err := c.walletClient.GetWalletsByUserId(context.Background(), &wallet_pb.UserId{
			UserId: fromWalletId,
		})
		if err != nil {
			return nil, types.ExtractGRPCErrDetails(err)
		}

		wallets, rerr := toWalletsProto(userWallets)
		if rerr != nil {
			return nil, rerr
		}

		if len(wallets) > 0 {
			for _, w := range wallets {
				if w.CoinId == 1 {
					// wlt, err := c.walletClient.GetWalletsByUserId(context.Background(), &wallet_pb.UserId{
					// 	UserId: 0,
					// })
					// if err != nil {
					// 	return nil, types.ExtractGRPCErrDetails(err)
					// }
					// if len(wlt.Wallets) <= 0 {
					// 	return nil, types.NewBadRequestError("Wallet Not found")
					// }
					c.walletClient.AddTransaction(context.Background(), &wallet_pb.AddTransactionRequest{
						Amount:           total,
						FromWalletId:     constants.MAIN_LUNC_WALLET_ID,
						FromWalletUserId: constants.MAIN_LUNC_USER_WALLET_ID,
						CoinId:           1,
						ToWalletAddress:  w.Address,
					})
					// JOB_QUEUE.Enqueue(Job{
					// 	Type: TransactionJob,
					// 	Transaction: models.AddTransactionDTO{
					// 		Amount:           total,
					// 		FromWalletId:     wlt.GetWallets()[0].Id,
					// 		FromWalletUserId: 0,
					// 		CoinId:           1,
					// 		ToWalletAddress:  w.Address,
					// 	},
					// })
					break
				}
			}
		}
		return nil, types.ExtractGRPCErrDetails(err)
	}
	if res != nil {
		return toTicketsProtos(res)
	} else {
		return nil, nil
	}
}
