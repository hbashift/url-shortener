package main

import (
	"github.com/hbashift/url-shortener/internal/api"
	"github.com/hbashift/url-shortener/internal/domain/repository/postgresDB"
	"github.com/hbashift/url-shortener/internal/domain/repository/redis"
	"github.com/hbashift/url-shortener/internal/server"
	"github.com/hbashift/url-shortener/internal/service"
	pb "github.com/hbashift/url-shortener/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

func RunGrpcServerWithRedis(redisConfig redis.Config, config1, config2 string) {
	rep := redis.NewRedis(&redisConfig)

	lis, err := net.Listen(config1, config2)
	if err != nil {
		log.Fatalf("failed to listen grpc server: %v", err)
	}

	s := service.NewShortenerService(rep)
	serv := server.NewShortenerServer(*s)
	grpcServ := grpc.NewServer()

	pb.RegisterShortenerServer(grpcServ, serv)

	err = grpcServ.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve grpc server: %v", err)
	}
}

func RunGrpcServerWithPg(pgConfig postgresDB.Config, c, config2 string, withGateway bool) {
	rep := postgresDB.NewPostgresDB(&pgConfig)

	lis, err := net.Listen(c, config2)
	if err != nil {
		log.Fatalf("failed to listen grpc server: %v", err)
	}

	s := service.NewShortenerService(rep)
	serv := server.NewShortenerServer(*s)
	grpcServ := grpc.NewServer()

	pb.RegisterShortenerServer(grpcServ, serv)

	if withGateway {
		time.Sleep(time.Second * 3)
		go api.RunHttpClient(":8080")
	}

	err = grpcServ.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve grpc server: %v", err)
	}
}

func main() {
	cfg := postgresDB.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "12345",
		DBName:   "shortener",
		SSLMode:  "disable",
	}

	RunGrpcServerWithPg(cfg, "tcp", ":8080", true)

	/*	cfg := redis.Config{
			Addr:        ":6379",
			Pass:        "0",
			DBNumMain:   0,
			DBNumUnique: 1,
		}

		RunGrpcServerWithRedis(cfg, "tcp", ":8080")*/
}
