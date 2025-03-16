package handlers

import (
	"context"
	"dolott_ticket/internal/models"
	"dolott_ticket/internal/services"
	pb "dolott_ticket/proto/api"
	"fmt"
	"strconv"
)

type TicketHandler struct {
	pb.UnimplementedTicketServiceServer
	ticketService services.TicketService
}

func NewTicketHandler(ticketService services.TicketService) *TicketHandler {
	return &TicketHandler{
		ticketService: ticketService,
	}
}

func (c *TicketHandler) AddTicket(ctx context.Context, request *pb.AddTicketRequest) (*pb.Ticket, error) {
	data := models.AddTicketDTO{
		UserId:     request.UserId,
		TicketType: request.TicketType,
	}

	res, err := c.ticketService.AddTicket(&data)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	return toTicketProto(res), nil
}

func (c *TicketHandler) GetTicketBySignatureAndUserId(ctx context.Context, request *pb.SignatureAndUserId) (*pb.Ticket, error) {
	res, err := c.ticketService.GetTicketBySignatureAndUserId(request.Signature, request.UserId)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	return toTicketProto(res), nil
}

func (c *TicketHandler) GetTicketByUserIdAndTicketId(ctx context.Context, request *pb.TicketIdAndUserId) (*pb.Ticket, error) {
	res, err := c.ticketService.GetTicketByUserIdAndTicketId(request.TicketId, request.UserId)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	return toTicketProto(res), nil
}

func (c *TicketHandler) UseTickets(ctx context.Context, request *pb.UseTicketsRequest) (*pb.Tickets, error) {
	res, err := c.ticketService.UseTickets(request.UserId, request.TotalUsingTickets, request.GameId)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	return toTicketsProtos(&models.Tickets{Tickets: res}), nil
}

func (c *TicketHandler) GetAllUserTickets(ctx context.Context, request *pb.UserIdAndPagination) (*pb.Tickets, error) {
	data := models.Pagination{
		Offset: request.Pagination.Offset,
		Limit:  request.Pagination.Limit,
		Total:  request.Pagination.GetTotal,
	}

	res, err := c.ticketService.GetAllUserTickets(request.UserId, &data)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	return toTicketsProtos(res), nil
}

func (c *TicketHandler) GetUserOpenTickets(ctx context.Context, request *pb.UserIdAndPagination) (*pb.Tickets, error) {
	data := models.Pagination{
		Offset: request.Pagination.Offset,
		Limit:  request.Pagination.Limit,
		Total:  request.Pagination.GetTotal,
	}

	res, err := c.ticketService.GetUserOpenTickets(request.UserId, &data)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	return toTicketsProtos(res), nil
}

func (c *TicketHandler) GetAllUsedTickets(ctx context.Context, request *pb.UserIdAndPagination) (*pb.Tickets, error) {
	data := models.Pagination{
		Offset: request.Pagination.Offset,
		Limit:  request.Pagination.Limit,
		Total:  request.Pagination.GetTotal,
	}

	res, err := c.ticketService.GetAllUsedTickets(request.UserId, &data)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	return toTicketsProtos(res), nil
}

func (c *TicketHandler) AddTickets(ctx context.Context, request *pb.AddTicketsRequest) (*pb.Tickets, error) {
	tickets := make([]*models.AddTicketDTO, 0)
	for _, data := range request.Tickets {
		tickets = append(tickets, &models.AddTicketDTO{
			UserId:     data.UserId,
			TicketType: data.TicketType,
		})
	}
	res, err := c.ticketService.AddTickets(tickets, request.ShouldReturn)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	if request.ShouldReturn {
		return toTicketsProtos(&models.Tickets{Tickets: res}), nil
	} else {
		return nil, nil
	}
}

func (c *TicketHandler) GetAllUserTicketsByGameId(ctx context.Context, request *pb.GetAllUserTicketsByGameIdRequest) (*pb.Tickets, error) {
	res, err := c.ticketService.GetAllUserTicketsByGameId(request.UserId, request.GameId)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	return toTicketsProtos(&models.Tickets{Tickets: res}), nil
}

func (c *TicketHandler) GetAllTicketsByGameId(ctx context.Context, request *pb.GameId) (*pb.Tickets, error) {
	res, err := c.ticketService.GetAllTicketsByGameId(request.GameId)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	return toTicketsProtos(&models.Tickets{Tickets: res}), nil
}

func (c *TicketHandler) GetAllPurchasedTicketsCountByGameId(ctx context.Context, request *pb.GameId) (*pb.Count, error) {
	res, err := c.ticketService.GetAllPurchasedTicketsCountByGameId(request.GameId)
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	return &pb.Count{Count: res}, nil
}

func toTicketProto(data *models.Ticket) *pb.Ticket {
	var usedAt *string
	if data.UsedAt != nil {
		fmt.Println(data.UsedAt)
		fmt.Println(data.UsedAt.Unix())
		fmt.Println(fmt.Sprintf("%d", data.UsedAt.Unix()))
		usedAtS := fmt.Sprintf("%d", data.UsedAt.Unix())
		usedAt = &usedAtS
	}

	return &pb.Ticket{
		Id:         data.ID,
		Signature:  data.Signature,
		UserId:     data.UserId,
		TicketType: data.TicketType,
		Status:     data.Status,
		Used:       data.Used,
		GameId:     data.GameId,
		UsedAt:     usedAt,
		CreatedAt:  strconv.FormatInt(data.CreatedAt.Unix(), 10),
	}
}

func toTicketsProtos(data *models.Tickets) *pb.Tickets {
	var tickets []*pb.Ticket
	for _, v := range data.Tickets {
		tickets = append(tickets, toTicketProto(&v))
	}
	return &pb.Tickets{Tickets: tickets, Total: data.Total}
}
