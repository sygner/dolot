package services

import (
	"dolott_ticket/internal/models"
	"dolott_ticket/internal/repository"
	"dolott_ticket/internal/types"
	"safir/libs/idgen"
)

type (
	TicketService interface {
		AddTicket(*models.AddTicketDTO) (*models.Ticket, *types.Error)
		GetTicketBySignatureAndUserId(string, int32) (*models.Ticket, *types.Error)
		GetTicketByUserIdAndTicketId(int32, int32) (*models.Ticket, *types.Error)
		GetAllUserTickets(int32, *models.Pagination) (*models.Tickets, *types.Error)
		GetUserOpenTickets(int32, *models.Pagination) (*models.Tickets, *types.Error)
		GetAllUsedTickets(int32, *models.Pagination) (*models.Tickets, *types.Error)
		UseTickets(int32, int32, string) ([]models.Ticket, *types.Error)
		AddTickets([]*models.AddTicketDTO, bool) ([]models.Ticket, *types.Error)
		GetAllUserTicketsByGameId(int32, string) ([]models.Ticket, *types.Error)
		GetAllTicketsByGameId(string) ([]models.Ticket, *types.Error)
		GetAllPurchasedTicketsCountByGameId(string) (int32, *types.Error)
	}
	ticketService struct {
		repository repository.TicketRepository
	}
)

func NewTicketService(ticketRepository repository.TicketRepository) TicketService {
	return &ticketService{repository: ticketRepository}
}

func (c *ticketService) AddTicket(data *models.AddTicketDTO) (*models.Ticket, *types.Error) {
	signature, err := idgen.NextNumericString(82)
	if err != nil {
		return nil, types.NewInternalError("Failed to generate ticket signature #6101")
	}

	data.Signature = signature
	data.Status = "open"

	res, rerr := c.repository.AddTicket(data)
	if rerr != nil {
		return nil, rerr
	}
	return res, nil
}

func (c *ticketService) GetTicketBySignatureAndUserId(signature string, userId int32) (*models.Ticket, *types.Error) {
	return c.repository.GetTicketBySignatureAndUserId(signature, userId)
}

func (c *ticketService) GetTicketByUserIdAndTicketId(ticketId int32, userId int32) (*models.Ticket, *types.Error) {
	return c.repository.GetTicketByUserIdAndTicketId(ticketId, userId)
}

func (c *ticketService) GetAllUserTickets(userId int32, pagInation *models.Pagination) (*models.Tickets, *types.Error) {
	res, err := c.repository.GetAllUserTickets(userId, pagInation)
	if err != nil {
		return nil, err
	}
	tickets := models.Tickets{Tickets: res}
	if pagInation.Total {
		res, err := c.repository.GetAllUserTicketsCount(userId)
		if err != nil {
			return nil, err
		}
		tickets.Total = &res
	}

	return &tickets, nil
}

func (c *ticketService) GetUserOpenTickets(userId int32, pagInation *models.Pagination) (*models.Tickets, *types.Error) {
	res, err := c.repository.GetUserOpenTickets(userId, pagInation)
	if err != nil {
		return nil, err
	}
	tickets := models.Tickets{Tickets: res}
	if pagInation.Total {
		res, err := c.repository.GetUserOpenTicketsCount(userId)
		if err != nil {
			return nil, err
		}
		tickets.Total = &res
	}

	return &tickets, nil
}

func (c *ticketService) GetAllUsedTickets(userId int32, pagInation *models.Pagination) (*models.Tickets, *types.Error) {
	res, err := c.repository.GetAllUsedTickets(userId, pagInation)
	if err != nil {
		return nil, err
	}
	tickets := models.Tickets{Tickets: res}
	if pagInation.Total {
		res, err := c.repository.GetAllUsedTicketsCount(userId)
		if err != nil {
			return nil, err
		}
		tickets.Total = &res
	}

	return &tickets, nil
}

func (c *ticketService) UseTickets(userId int32, totalUsingTicket int32, gameId string) ([]models.Ticket, *types.Error) {
	return c.repository.UseTickets(userId, totalUsingTicket, gameId)
}

func (c *ticketService) AddTickets(data []*models.AddTicketDTO, shouldReturn bool) ([]models.Ticket, *types.Error) {
	for _, ticket := range data {
		signature, err := idgen.NextNumericString(82)
		if err != nil {
			return nil, types.NewInternalError("Failed to generate ticket signature #6102")
		}

		ticket.Signature = signature
		ticket.Status = "open"
	}
	res, err := c.repository.AddTickets(data, shouldReturn)
	if err != nil {
		return nil, err
	}
	if shouldReturn {
		return res, nil
	} else {
		return nil, nil
	}
}

func (c *ticketService) GetAllUserTicketsByGameId(userId int32, gameId string) ([]models.Ticket, *types.Error) {
	return c.repository.GetAllUserTicketsByGameId(userId, gameId)
}

func (c *ticketService) GetAllTicketsByGameId(gameId string) ([]models.Ticket, *types.Error) {
	return c.repository.GetAllTicketsByGameId(gameId)
}

func (c *ticketService) GetAllPurchasedTicketsCountByGameId(gameId string) (int32, *types.Error) {
	return c.repository.GetAllPurchasedTicketsCountByGameId(gameId)
}
