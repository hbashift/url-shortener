package main

import (
	"github.com/hbashift/url-shortener/internal/api"
)

func main() {
	api.RunHttpClient(":8080")
}
