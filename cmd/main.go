package main

import (
	"go-base/config"
	"go-base/internal/infrastructure/initializer"
)

func main() {
	cfg := config.Load()

	initializer.InitializeServer(cfg)
}
