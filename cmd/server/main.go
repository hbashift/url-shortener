package main

import (
	"github.com/hbashift/url-shortener/internal/domain/repository/postgresDB"
	"github.com/hbashift/url-shortener/internal/domain/repository/redis"
	"github.com/hbashift/url-shortener/internal/server"
	"github.com/hbashift/url-shortener/internal/service"
	pb "github.com/hbashift/url-shortener/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"strconv"
)

func RunGrpcServerWithRedis(redisConfig redis.Config, protocol, port string) {
	rep := redis.NewRedis(&redisConfig)

	lis, err := net.Listen(protocol, port)
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

func RunGrpcServerWithPg(pgConfig postgresDB.Config, protocol, port string) {
	rep := postgresDB.NewPostgresDB(&pgConfig)

	lis, err := net.Listen(protocol, port)
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

func main() {
	dbType := os.Getenv("DB_TYPE")

	if dbType == "postgres" {

		cfg := postgresDB.Config{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			Username: os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			DBName:   os.Getenv("POSTGRES_DB"),
			SSLMode:  os.Getenv("POSTGRES_SSLMODE"),
		}

		RunGrpcServerWithPg(cfg, "tcp", os.Getenv("PORT"))

	} else if dbType == "redis" {
		mainDB, err := strconv.Atoi(os.Getenv("MAIN_DB"))
		if err != nil {
			panic(err)
		}

		uniqueDB, err := strconv.Atoi(os.Getenv("UNIQUE_DB"))
		if err != nil {
			panic(err)
		}

		cfg := redis.Config{
			Addr:        os.Getenv("REDIS_ADDR"),
			Pass:        os.Getenv("PASS"),
			DBNumMain:   mainDB,
			DBNumUnique: uniqueDB,
		}

		RunGrpcServerWithRedis(cfg, "tcp", os.Getenv("PORT"))
	}
}
