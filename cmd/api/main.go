package main

import (
	"github.com/hbashift/url-shortener/internal/api"
)

func main() {

	api.RunHttpClient("server"+":8080", ":9090")
}
