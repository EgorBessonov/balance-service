//Package server represents balance gRPC server handlers
package server

import (
	"context"
	"fmt"
	"github.com/EgorBessonov/balance-service/internal/service"
	balanceService2 "github.com/EgorBessonov/balance-service/protocol"
)

//Server type represents server structure in balance service
type Server struct {
	service *service.Service
	balanceService2.UnimplementedBalanceServer
}

//NewServer returns new server instance
func NewServer(service *service.Service) *Server {
	return &Server{
		service: service,
	}
}

//Get method returns user balance
func (server *Server) Get(ctx context.Context, request *balanceService2.GetRequest) (*balanceService2.GetResponse, error) {
	balance, err := server.service.Get(ctx, request.UserId)
	if err != nil {
		return nil, err
	}
	return &balanceService2.GetResponse{Balance: balance}, nil
}

//Check method checks if user has necessary balance
func (server *Server) Check(ctx context.Context, request *balanceService2.CheckRequest) (*balanceService2.CheckResponse, error) {
	ok, err := server.service.Check(ctx, request.UserId, request.RequiredBalance)
	if err != nil {
		return nil, err
	}
	return &balanceService2.CheckResponse{Ok: ok}, nil
}

//TopUp method increase user balance
func (server *Server) TopUp(ctx context.Context, request *balanceService2.TopUpRequest) (*balanceService2.TopUpResponse, error) {
	err := server.service.TopUp(ctx, request.UserId, request.Shift)
	if err != nil {
		return nil, err
	}
	return &balanceService2.TopUpResponse{Result: fmt.Sprint("success")}, nil
}

//Withdraw method decrease user balance
func (server *Server) Withdraw(ctx context.Context, request *balanceService2.WithdrawRequest) (*balanceService2.WithdrawResponse, error) {
	err := server.service.Withdraw(ctx, request.UserId, request.Shift)
	if err != nil {
		return nil, err
	}
	return &balanceService2.WithdrawResponse{Result: "success"}, nil
}
