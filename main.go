package main

import (
	"fmt"
	"github.com/hbashift/url-shortener/internal/util/encoder"
)

func main() {
	/*cfgPg := postgres.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "12345",
		DBName:   "shortener",
		SSLMode:  "disable",
	}
	rep := postgres.NewPostgresDB(&cfgPg)

		cfgRedis := redis.Config{
			Addr:        ":6379",
			Pass:        "",
			DBNumMain:   0,
			DBNumUnique: 1,
		}
		rep := redis.NewRedis(&cfgRedis)

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
	}*/

	fmt.Println(encoder.DecryptUrl("CdFksa_sdf"))
}
