package repository

import (
	"database/sql"
	"dolott_ticket/internal/models"
	"dolott_ticket/internal/types"
	"fmt"
	"strings"

	"github.com/lib/pq"
)

type (
	TicketRepository interface {
		AddTicket(*models.AddTicketDTO) (*models.Ticket, *types.Error)
		GetTicketBySignatureAndUserId(string, int32) (*models.Ticket, *types.Error)
		GetTicketByUserIdAndTicketId(int32, int32) (*models.Ticket, *types.Error)
		GetAllUserTickets(int32, *models.Pagination) ([]models.Ticket, *types.Error)
		GetUserOpenTickets(int32, *models.Pagination) ([]models.Ticket, *types.Error)
		GetAllUsedTickets(int32, *models.Pagination) ([]models.Ticket, *types.Error)
		AddTickets([]*models.AddTicketDTO, bool) ([]models.Ticket, *types.Error)
		GetAllPurchasedTicketsCountByGameId(string) (int32, *types.Error)

		GetAllUserTicketsCount(int32) (int32, *types.Error)
		GetUserOpenTicketsCount(int32) (int32, *types.Error)
		GetAllUsedTicketsCount(int32) (int32, *types.Error)

		UseTickets(int32, int32, string) ([]models.Ticket, *types.Error)
		GetAllUserTicketsByGameId(int32, string) ([]models.Ticket, *types.Error)
		GetAllTicketsByGameId(string) ([]models.Ticket, *types.Error)
	}
	ticketRepository struct {
		db *sql.DB
	}
)

func NewTicketRepository(db *sql.DB) TicketRepository {
	return &ticketRepository{
		db: db,
	}
}

func (c *ticketRepository) AddTicket(data *models.AddTicketDTO) (*models.Ticket, *types.Error) {
	query := `INSERT INTO tickets (signature, user_id, ticket_type, status) VALUES ($1, $2, $3, $4) RETURNING id, signature, user_id, ticket_type, status, used, game_id, used_at, created_at`
	row := c.db.QueryRow(query, data.Signature, data.UserId, data.TicketType, data.Status)
	if row.Err() != nil {
		return nil, types.NewInternalError("Failed to add ticket, error code #6001")
	}
	ticket := &models.Ticket{}
	err := row.Scan(&ticket.ID, &ticket.Signature, &ticket.UserId, &ticket.TicketType, &ticket.Status, &ticket.Used, &ticket.GameId, &ticket.UsedAt, &ticket.CreatedAt)
	if err != nil {
		return nil, types.NewInternalError("Failed to fetch ticket, error code #6002")
	}
	return ticket, nil
}

func (c *ticketRepository) GetTicketBySignatureAndUserId(signature string, userId int32) (*models.Ticket, *types.Error) {
	query := `SELECT id, signature, user_id, ticket_type, status, used, game_id, used_at, created_at FROM tickets WHERE signature = $1 AND user_id = $2`

	ticket := &models.Ticket{}
	err := c.db.QueryRow(query, signature, userId).Scan(&ticket.ID, &ticket.Signature, &ticket.UserId, &ticket.TicketType, &ticket.Status, &ticket.Used, &ticket.UsedAt, &ticket.GameId, &ticket.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("ticket not found, error code #6003")
		}
		return nil, types.NewInternalError("failed to get ticket, error code #6004")
	}

	return ticket, nil
}

func (c *ticketRepository) GetTicketByUserIdAndTicketId(ticketId int32, userId int32) (*models.Ticket, *types.Error) {
	query := `SELECT id, signature, user_id, ticket_type, status, used, game_id, used_at, created_at FROM tickets WHERE id = $1 AND user_id = $2`
	ticket := &models.Ticket{}
	err := c.db.QueryRow(query, ticketId, userId).Scan(&ticket.ID, &ticket.Signature, &ticket.UserId, &ticket.TicketType, &ticket.Status, &ticket.Used, &ticket.UsedAt, &ticket.GameId, &ticket.CreatedAt)

	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("ticket not found, error code #6005")
		}
		return nil, types.NewInternalError("failed to get ticket, error code #6006")
	}

	return ticket, nil
}

func (c *ticketRepository) GetAllUserTickets(userId int32, pagInation *models.Pagination) ([]models.Ticket, *types.Error) {
	query := `SELECT id, signature, user_id, ticket_type, status, used, game_id, used_at, created_at FROM tickets WHERE user_id = $1 ORDER BY created_at DESC OFFSET $2 LIMIT $3`
	rows, err := c.db.Query(query, userId, pagInation.Offset, pagInation.Limit)

	if err != nil {
		return nil, types.NewInternalError("failed to get ticket, error code #6007")
	}

	tickets := make([]models.Ticket, 0)

	for rows.Next() {
		ticket := &models.Ticket{}
		err := rows.Scan(&ticket.ID, &ticket.Signature, &ticket.UserId, &ticket.TicketType, &ticket.Status, &ticket.Used, &ticket.UsedAt, &ticket.GameId, &ticket.CreatedAt)

		if err != nil {
			return nil, types.NewInternalError("failed to fetch ticket, error code #6008")
		}
		tickets = append(tickets, *ticket)
	}

	return tickets, nil
}

func (c *ticketRepository) GetUserOpenTickets(userId int32, pagInation *models.Pagination) ([]models.Ticket, *types.Error) {
	query := `SELECT id, signature, user_id, ticket_type, status, used, game_id, used_at, created_at FROM tickets WHERE user_id = $1 AND used = false ORDER BY created_at DESC OFFSET $2 LIMIT $3`
	rows, err := c.db.Query(query, userId, pagInation.Offset, pagInation.Limit)

	if err != nil {
		return nil, types.NewInternalError("failed to get ticket, error code #6009")
	}

	tickets := make([]models.Ticket, 0)

	for rows.Next() {
		ticket := &models.Ticket{}
		err := rows.Scan(&ticket.ID, &ticket.Signature, &ticket.UserId, &ticket.TicketType, &ticket.Status, &ticket.Used, &ticket.UsedAt, &ticket.GameId, &ticket.CreatedAt)

		if err != nil {
			return nil, types.NewInternalError("failed to fetch ticket, error code #6010")
		}
		tickets = append(tickets, *ticket)
	}

	return tickets, nil
}

func (c *ticketRepository) GetAllUsedTickets(userId int32, pagInation *models.Pagination) ([]models.Ticket, *types.Error) {
	query := `SELECT id, signature, user_id, ticket_type, status, used, game_id, used_at, created_at FROM tickets WHERE user_id = $1 AND used = true ORDER BY created_at DESC OFFSET $2 LIMIT $3`
	rows, err := c.db.Query(query, userId, pagInation.Offset, pagInation.Limit)

	if err != nil {
		return nil, types.NewInternalError("failed to get ticket, error code #6011")
	}

	tickets := make([]models.Ticket, 0)

	for rows.Next() {
		ticket := &models.Ticket{}
		err := rows.Scan(&ticket.ID, &ticket.Signature, &ticket.UserId, &ticket.TicketType, &ticket.Status, &ticket.Used, &ticket.GameId, &ticket.UsedAt, &ticket.CreatedAt)

		if err != nil {
			return nil, types.NewInternalError("failed to fetch ticket, error code #6012")
		}
		tickets = append(tickets, *ticket)
	}

	return tickets, nil
}

func (c *ticketRepository) GetAllPurchasedTicketsCountByGameId(gameId string) (int32, *types.Error) {
	query := `SELECT count(*) FROM "tickets" WHERE game_id = $1 AND used = true`
	var totalCount int32
	err := c.db.QueryRow(query, gameId).Scan(&totalCount)
	if err != nil {
		fmt.Println(err)
		return 0, types.NewInternalError("internal issue, error code #6037")
	}
	return totalCount, nil
}

func (c *ticketRepository) GetAllUserTicketsCount(userid int32) (int32, *types.Error) {
	query := `SELECT count(*) FROM "tickets" WHERE user_id = $1 AND used = false`
	var totalCount int32
	err := c.db.QueryRow(query, userid).Scan(&totalCount)
	if err != nil {
		fmt.Println(err)
		return 0, types.NewInternalError("internal issue, error code #6013")
	}
	return totalCount, nil
}

func (c *ticketRepository) GetUserOpenTicketsCount(userid int32) (int32, *types.Error) {
	query := `SELECT count(*) FROM "tickets" WHERE user_id = $1 AND used = false`
	var totalCount int32
	err := c.db.QueryRow(query, userid).Scan(&totalCount)
	if err != nil {
		fmt.Println(err)
		return 0, types.NewInternalError("internal issue, error code #6014")
	}
	return totalCount, nil
}

func (c *ticketRepository) GetAllUsedTicketsCount(userid int32) (int32, *types.Error) {
	query := `SELECT count(*) FROM "tickets" WHERE user_id = $1 AND used = true`
	var totalCount int32
	err := c.db.QueryRow(query, userid).Scan(&totalCount)
	if err != nil {
		fmt.Println(err)
		return 0, types.NewInternalError("internal issue, error code #6015")
	}
	return totalCount, nil
}

func (c *ticketRepository) UseTickets(userId int32, totalUsingTicket int32, gameId string) ([]models.Ticket, *types.Error) {
	tx, err := c.db.Begin()
	if err != nil {
		return nil, types.NewInternalError("Failed to start transaction, error code #6020")
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var availableCount int32
	countQuery := `SELECT COUNT(*) FROM tickets WHERE user_id = $1 AND used = false`
	err = tx.QueryRow(countQuery, userId).Scan(&availableCount)
	if err != nil {
		return nil, types.NewInternalError("Failed to check available tickets, error code #6021")
	}

	if availableCount < totalUsingTicket {
		return nil, types.NewBadRequestError("Not enough tickets available for that amount #6021-2")
	}

	query := `SELECT id, signature, user_id, ticket_type, status, used, game_id, used_at, created_at 
              FROM tickets WHERE user_id = $1 AND used = false 
              ORDER BY created_at ASC LIMIT $2 FOR UPDATE`
	rows, err := tx.Query(query, userId, totalUsingTicket)
	if err != nil {
		return nil, types.NewInternalError("Failed to fetch tickets, error code #6022")
	}
	defer rows.Close()

	tickets := []models.Ticket{}
	ticketIds := []int32{}

	for rows.Next() {
		ticket := models.Ticket{}
		err := rows.Scan(&ticket.ID, &ticket.Signature, &ticket.UserId, &ticket.TicketType, &ticket.Status, &ticket.Used, &ticket.GameId, &ticket.UsedAt, &ticket.CreatedAt)
		if err != nil {
			return nil, types.NewInternalError("Failed to process ticket data, error code #6023")
		}
		tickets = append(tickets, ticket)
		ticketIds = append(ticketIds, ticket.ID)
	}

	if len(tickets) == 0 {
		return nil, types.NewBadRequestError("Not enough tickets available #6023-2")
	}
	updateQuery := `UPDATE tickets SET used = true, used_at = NOW(), game_id = $1 WHERE id = ANY($2)`
	_, err = tx.Exec(updateQuery, gameId, pq.Array(ticketIds))
	if err != nil {
		return nil, types.NewInternalError("Failed to update tickets as used, error code #6024")
	}

	err = tx.Commit()
	if err != nil {
		return nil, types.NewInternalError("Failed to commit transaction, error code #6025")
	}
	return tickets, nil
}

func (c *ticketRepository) AddTickets(dataList []*models.AddTicketDTO, shouldReturn bool) ([]models.Ticket, *types.Error) {
	if len(dataList) == 0 {
		return nil, types.NewBadRequestError("No tickets to add")
	}

	tx, err := c.db.Begin()
	if err != nil {
		return nil, types.NewInternalError("Failed to start transaction, error code #6029")
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	var valueStrings []string
	var valueArgs []interface{}
	for i, data := range dataList {
		fmt.Println("Data", *data)
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d)", i*4+1, i*4+2, i*4+3, i*4+4))
		valueArgs = append(valueArgs, data.Signature, data.UserId, data.TicketType, data.Status)
	}

	query := fmt.Sprintf(`INSERT INTO tickets (signature, user_id, ticket_type, status) 
                          VALUES %s RETURNING id, signature, user_id, ticket_type, status, used, game_id, used_at, created_at`,
		strings.Join(valueStrings, ","))
	fmt.Println(query)
	if !shouldReturn {
		_, err = tx.Exec(query, valueArgs...)
		if err != nil {
			return nil, types.NewInternalError("Failed to insert tickets, error code #6030-2")
		}
		err = tx.Commit()
		if err != nil {
			return nil, types.NewInternalError("Failed to commit transaction, error code #6032-2")
		}
		return nil, nil
	} else {

		rows, err := tx.Query(query, valueArgs...)
		if err != nil {
			return nil, types.NewInternalError("Failed to insert tickets, error code #6030")
		}
		defer rows.Close()
		fmt.Println("DONE BUYING TICKETS")
		var addedTickets []models.Ticket
		for rows.Next() {
			ticket := models.Ticket{}
			err := rows.Scan(&ticket.ID, &ticket.Signature, &ticket.UserId, &ticket.TicketType, &ticket.Status, &ticket.Used, &ticket.GameId, &ticket.UsedAt, &ticket.CreatedAt)
			if err != nil {
				return nil, types.NewInternalError("Failed to process inserted tickets, error code #6031")
			}
			fmt.Println("TICKET")
			addedTickets = append(addedTickets, ticket)
		}

		err = tx.Commit()
		if err != nil {
			return nil, types.NewInternalError("Failed to commit transaction, error code #6032")
		}
		fmt.Println("ADDED TICKETs", addedTickets)
		return addedTickets, nil
	}
}

func (c *ticketRepository) GetAllUserTicketsByGameId(userId int32, gameId string) ([]models.Ticket, *types.Error) {
	query := `SELECT id, signature, user_id, ticket_type, status, used, used_at, game_id, created_at FROM tickets WHERE user_id = $1 AND game_id = $2`
	rows, err := c.db.Query(query, userId, gameId)

	if err != nil {
		return nil, types.NewInternalError("failed to get ticket, error code #6033")
	}

	tickets := make([]models.Ticket, 0)

	for rows.Next() {
		ticket := &models.Ticket{}
		err := rows.Scan(&ticket.ID, &ticket.Signature, &ticket.UserId, &ticket.TicketType, &ticket.Status, &ticket.Used, &ticket.UsedAt, &ticket.GameId, &ticket.CreatedAt)

		if err != nil {
			fmt.Println(err)
			return nil, types.NewInternalError("failed to fetch ticket, error code #6034")
		}
		tickets = append(tickets, *ticket)
	}

	return tickets, nil
}

func (c *ticketRepository) GetAllTicketsByGameId(gameId string) ([]models.Ticket, *types.Error) {
	query := `SELECT id, signature, user_id, ticket_type, status, used, used_at, game_id, created_at FROM tickets WHERE game_id = $1`
	rows, err := c.db.Query(query, gameId)

	if err != nil {
		return nil, types.NewInternalError("failed to get ticket, error code #6035")
	}

	tickets := make([]models.Ticket, 0)

	for rows.Next() {
		ticket := &models.Ticket{}
		err := rows.Scan(&ticket.ID, &ticket.Signature, &ticket.UserId, &ticket.TicketType, &ticket.Status, &ticket.Used, &ticket.UsedAt, &ticket.GameId, &ticket.CreatedAt)
		if err != nil {
			fmt.Println(err)
			return nil, types.NewInternalError("failed to fetch ticket, error code #6036")
		}
		tickets = append(tickets, *ticket)
	}

	return tickets, nil
}
