package main

import (
	"context"
	"github.com/EgorBessonov/balance-service/internal/config"
	"github.com/EgorBessonov/balance-service/internal/repository"
	"github.com/EgorBessonov/balance-service/internal/server"
	"github.com/EgorBessonov/balance-service/internal/service"
	"github.com/EgorBessonov/balance-service/protocol"
	"github.com/caarlos0/env"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("balance service: can't parse config")
	}
	postgres, err := pgxpool.Connect(context.Background(), cfg.PostgresURL)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("balance service: can't connect to the postgres database")
	}
	rps := repository.NewPostgresRepository(postgres)
	serv := service.NewService(rps)
	balanceServer := server.NewServer(serv)
	newgRPCServer(cfg.BalanceServerPort, balanceServer)
}

func newgRPCServer(port string, s *server.Server) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("balance service: can't create gRPC server")
	}
	gServer := grpc.NewServer()
	balanceService.RegisterBalanceServer(gServer, s)
	log.Printf("price service: listnening at %s", lis.Addr())
	if err = gServer.Serve(lis); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("balance service: gRPC server failed")
	}
}
