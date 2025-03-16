package services

import (
	"dolott_user_gw_http/internal/models"
	"dolott_user_gw_http/internal/types"
	"dolott_user_gw_http/internal/utils"
	game_pb "dolott_user_gw_http/proto/api/game"
	profile_pb "dolott_user_gw_http/proto/api/profile"
	ticket_pb "dolott_user_gw_http/proto/api/ticket"
	wallet_pb "dolott_user_gw_http/proto/api/wallet"
	"fmt"
	"log"
	"time"
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
		Division:    res.Division,
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

func toWinnerProto(res *game_pb.Winner) (*models.Winners, *types.Error) {
	createdAt, err := utils.ParseTime(res.CreatedAt, "failed to convert the created at, wrong format #1-8")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &models.Winners{
		Id:           res.Id,
		GameId:       res.GameId,
		GameType:     res.GameType,
		Divisions:    toDivisionResultsProto(res.Divisions.DivisionResults),
		ResultNumber: res.ResultNumber,
		Prize:        res.Prize,
		JackPot:      res.Jackpot,
		TotalPaid:    res.TotalPaid,
		CreatedAt:    createdAt,
	}, nil
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

func toCoinProto(res *wallet_pb.Coin) *models.Coin {
	return &models.Coin{
		CoinId: res.CoinId,
		Name:   res.CurrencyName,
		Symbol: res.CurrencySymbol,
	}
}

func toCoinsProto(res *wallet_pb.Coins) []models.Coin {
	coins := make([]models.Coin, 0)
	for _, coin := range res.Coins {
		coins = append(coins, *toCoinProto(coin))
	}
	return coins
}

func toWalletProto(res *wallet_pb.Wallet) (*models.Wallet, *types.Error) {
	createdAt, err := utils.ParseTime(res.CreatedAt, "failed to convert the created at, wrong format #1-1")
	if err != nil {
		return nil, err
	}
	updatedAt, err := utils.ParseTime(res.UpdatedAt, "failed to convert the updated at, wrong format #1-2")
	if err != nil {
		return nil, err
	}
	return &models.Wallet{
		Id:        res.Id,
		Sid:       res.Sid,
		UserId:    res.UserId,
		CoinId:    res.CoinId,
		Balance:   float64(res.Balance),
		Address:   res.Address,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func toWalletsProto(res *wallet_pb.Wallets) ([]models.Wallet, *types.Error) {
	wallets := make([]models.Wallet, 0)
	for _, wallet := range res.Wallets {
		w, err := toWalletProto(wallet)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, *w)
	}
	return wallets, nil
}

func toTransactionProto(res *wallet_pb.Transaction) (*models.Transaction, *types.Error) {
	transactionAt, err := utils.ParseTime(res.TransactionAt, "failed to convert the transaction at, wrong format #1-3")
	if err != nil {
		return nil, err
	}
	return &models.Transaction{
		TxId:          res.TxId,
		CurrencyId:    res.CurrencyId,
		CurrencyName:  res.CurrencyName,
		FromAddress:   res.FromAddress,
		ToAddress:     res.ToAddress,
		FromWalletId:  res.FromWalletId,
		FromPublicKey: res.FromPublicKey,
		Amount:        float64(res.Amount),
		TransactionAt: transactionAt,
	}, nil
}

func toTransactionsProto(res *wallet_pb.Transactions) ([]models.Transaction, *types.Error) {
	transactions := make([]models.Transaction, 0)
	for _, transaction := range res.Transactions {
		t, err := toTransactionProto(transaction)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, *t)
	}
	return transactions, nil
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
		Prize:            res.Prize,
		AutoCompute:      res.AutoCompute,
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
		HighestRank:   res.HighestRank,
	}, nil
}

func toTicketProto(res *ticket_pb.Ticket) (*models.Ticket, *types.Error) {
	createdAt, err := utils.ParseTime(res.CreatedAt, "failed to convert the created at, wrong format #1-4")
	if err != nil {
		return nil, err
	}
	var usedAt *time.Time
	if res.UsedAt != nil {
		usedAtS, err := utils.ParseTime(*res.UsedAt, "failed to convert the created at, wrong format #1-4")
		if err != nil {
			return nil, err
		}
		usedAt = &usedAtS
	}
	return &models.Ticket{
		ID:         res.Id,
		Signature:  res.Signature,
		UserId:     res.UserId,
		TicketType: res.TicketType,
		Status:     res.Status,
		Used:       res.Used,
		UsedAt:     usedAt,
		GameId:     res.GameId,
		CreatedAt:  createdAt,
	}, nil
}

func toTicketsProtos(res *ticket_pb.Tickets) (*models.Tickets, *types.Error) {
	tickets := make([]models.Ticket, 0)
	for _, data := range res.Tickets {
		ticket, err := toTicketProto(data)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, *ticket)
	}
	return &models.Tickets{Tickets: tickets, Total: res.Total}, nil
}

var JOB_QUEUE JobQueue

type JobType int

type JobPriority int

const (
	LotteryJob JobType = iota
	TransactionJob
)

const (
	Force JobPriority = iota
	Important
	Normal
)

type Job struct {
	Type        JobType
	Priority    JobPriority
	GameID      string                   // Only for Lottery jobs
	Results     []models.DivisionResult  // Only for Lottery jobs
	Transaction models.AddTransactionDTO // Only for Transaction jobs
}

type JobQueue struct {
	normalJobs    chan Job
	priorityJobs  chan Job
	quit          chan struct{}
	gameService   GameService
	walletService WalletService
}

func NewJobQueue(gameService GameService, walletService WalletService) JobQueue {
	jq := &JobQueue{
		normalJobs:    make(chan Job, 100), // Buffered channel for normal jobs
		priorityJobs:  make(chan Job, 100), // Buffered channel for priority jobs
		quit:          make(chan struct{}),
		gameService:   gameService,
		walletService: walletService,
	}

	// Start the worker
	go jq.startWorker()
	return *jq
}

func (jq *JobQueue) startWorker() {
	for {
		select {
		case job := <-jq.priorityJobs:
			// Always process priority jobs first
			jq.processJob(job)
		case job := <-jq.normalJobs:
			// Only process normal jobs if no priority jobs are present
			jq.processJob(job)
		case <-jq.quit:
			log.Println("Job queue shutting down...")
			return
		}
	}
}

func (jq *JobQueue) processJob(job Job) {
	log.Printf("Processing Job: Type=%d, Priority=%d", job.Type, job.Priority)

	switch job.Type {
	case LotteryJob:
		log.Printf("Processing Lottery Job: GameID=%s", job.GameID)
		jq.gameService.TransactionForWinner(job.Results, job.GameID)

	case TransactionJob:
		log.Printf("Processing Transaction: From=%d To=%s Amount=%.2f",
			job.Transaction.FromWalletId, job.Transaction.ToWalletAddress, job.Transaction.Amount)

		res, err := jq.walletService.AddTransaction(&job.Transaction)
		if err != nil {
			log.Printf("Transaction failed: %v", err)
		}
		log.Println("Buying Ticket", res)
	}
}

// Enqueue a Job with Priority Handling
func (jq *JobQueue) Enqueue(job Job) {
	// If the job is Force or Important priority, enqueue to priorityJobs
	switch job.Priority {
	case Force, Important:
		select {
		case jq.priorityJobs <- job:
			log.Println("[JOB] Enqueued Important/Force Job")
		default:
			log.Println("[ERROR] Priority Queue is full. Unable to enqueue Important/Force Job.")
		}
	case Normal:
		// If the job is Normal priority, enqueue to normalJobs
		select {
		case jq.normalJobs <- job:
			log.Println("[JOB] Enqueued Normal Job")
		default:
			log.Println("[ERROR] Normal Queue is full. Unable to enqueue Normal Job.")
		}
	}
}

// Shutdown gracefully
func (jq *JobQueue) Shutdown() {
	close(jq.quit)
}

// Division 1: Jackpot
// Division 2: 11.3% of Jackpot
// Division 3: 17.1% of Jackpot
// Division 4: 25.34% of Jackpot
// Division 5: 37.4% of Jackpot
var LottoWinnerDivisions = map[string]float64{
	"Division 1": 1.0,    // 100% of Jackpot
	"Division 2": 0.374,  // 37.4% of Jackpot
	"Division 3": 0.2534, // 25.34% of Jackpot
	"Division 4": 0.171,  // 17.1% of Jackpot
	"Division 5": 0.113,  // 11.3% of Jackpot
}

// Division 2: 5.5% of the jackpot
// Division 3: 6.5% of the jackpot
// Division 4: 5.0% of the jackpot
// Division 5: 4.0% of the jackpot
// Division 6: 49.5% of the jackpot
var OsLottoWinnerDivisions = map[string]float64{
	"Division 1": 1.0,   // 100% of Jackpot
	"Division 2": 0.055, // 5.5% of Jackpot
	"Division 3": 0.065, // 6.5% of Jackpot
	"Division 4": 0.05,  // 5.0% of Jackpot
	"Division 5": 0.04,  // 4.0% of Jackpot
	"Division 6": 0.495, // 49.5% of Jackpot
}

// Division 1: Jackpot
// Division 2: 5.52% of the jackpot
// Division 3: 6.87% of the jackpot
// Division 4: 5.97% of the jackpot
// Division 5: 4.48% of the jackpot
// Division 6: 28.06% of the jackpot
var PowerballWinnerDivisions = map[string]float64{
	"Division 1": 1.0,    // 100% of Jackpot
	"Division 2": 0.0552, // 5.52% of Jackpot
	"Division 3": 0.0687, // 6.87% of Jackpot
	"Division 4": 0.0597, // 5.97% of Jackpot
	"Division 5": 0.0448, // 4.48% of Jackpot
	"Division 6": 0.2806, // 28.06% of Jackpot
}

// Division 1: Jackpot (increases with rollovers)
// Division 2: 5% of the starting jackpot (fixed prize amount for each winner, does not increase with rollovers)
// Division 3: 0.25% of the starting jackpot (fixed prize amount for each winner, does not increase with rollovers)
// Division 4: 0.0005% of the starting jackpot (fixed prize amount for each winner, does not increase with rollovers)
// Division 5: 0.0005% of the starting jackpot (fixed prize amount for each winner, does not increase with rollovers)
// Division 6: 0.000035% of the starting jackpot (fixed prize amount for each winner, does not increase with rollovers)
var AmericanPowerballWinnerDivisions = map[string]float64{
	"Division 1": 1.0,        // 100% of Jackpot
	"Division 2": 0.05,       // 5% of Jackpot
	"Division 3": 0.0025,     // 0.25% of Jackpot
	"Division 4": 0.000005,   // 0.0005% of Jackpot
	"Division 5": 0.000005,   // 0.0005% of Jackpot
	"Division 6": 0.00000035, // 0.000035% of Jackpot

}
