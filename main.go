package main

import (
	"github.com/hbashift/url-shortener/internal/domain/repository/postgres"
	"github.com/hbashift/url-shortener/internal/server"
	"github.com/hbashift/url-shortener/internal/service"
	pb "github.com/hbashift/url-shortener/pb"
	"google.golang.org/grpc"
	"net"
)

func main() {
	cfgPg := postgres.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "12345",
		DBName:   "shortener",
		SSLMode:  "disable",
	}
	rep := postgres.NewPostgresDB(&cfgPg)

	/*	cfgRedis := redis.Config{
			Addr:        ":6379",
			Pass:        "",
			DBNumMain:   0,
			DBNumUnique: 1,
		}
		rep := redis.NewRedis(&cfgRedis)*/

	s := service.NewShortenerService(rep)
	serv := server.NewShortenerServer(s)
	grpcServer := grpc.NewServer()
	pb.RegisterShortenerServer(grpcServer, serv)
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	if err = grpcServer.Serve(l); err != nil {
		panic(err)
	}
}
