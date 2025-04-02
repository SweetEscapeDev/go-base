package main

import (
	"go-base/config"
	"go-base/internal/infrastructure/database"
	"go-base/internal/infrastructure/handler/http"
)

func main() {
	c := config.Load()

	database.Init(c)

	http.Init(c)
}
