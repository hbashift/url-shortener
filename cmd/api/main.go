package main

import (
	"github.com/hbashift/url-shortener/internal/api"
	"os"
)

func main() {
	api.RunHttpClient(os.Getenv("GRPC_ADDR"), os.Getenv("GATEWAY_PORT"))
}
